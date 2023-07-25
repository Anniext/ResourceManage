package services

import (
	"ResourceManage/data"
	"ResourceManage/model"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type rangeBytes struct {
	start  int64
	end    int64
	length int64
}

func (r *RouterGroup) Upload(c *gin.Context) {
	var file model.AvtFile
	read := c.Request
	f, handler, err := read.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}(f)
	if err := data.UploadFile(handler, f, &file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, file)
}

// Download 断点续传、暂停和继续下载
func (r *RouterGroup) Download(c *gin.Context) {
	filePath := c.Query("path")
	if filePath == "" {
		c.String(http.StatusBadRequest, "File path is missing")
		return
	}

file, err := os.Open(filePath)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to open file")
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println("Error close file", err)
		}
	}(file)

	fileInfo, err := file.Stat()
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to get file info")
		return
	}

	fileSize := fileInfo.Size()
	rangeHeader := c.GetHeader("Range")

	if rangeHeader != "" {
		rangeBytes, err := parseRangeHeader(rangeHeader, fileSize)
		if err != nil {
			c.String(http.StatusRequestedRangeNotSatisfiable, "Invalid range header")
			return
		}

		c.Status(http.StatusPartialContent)
		c.Header("Accept-Ranges", "bytes")
		c.Header("Content-Range", fmt.Sprintf("bytes %d-%d/%d", rangeBytes.start, rangeBytes.end, fileSize))
		c.Header("Content-Length", fmt.Sprintf("%d", rangeBytes.length))
		file.Seek(rangeBytes.start, io.SeekStart)

		buf := make([]byte, 4096)
		for {
			n, err := file.Read(buf)
			if err != nil {
				if err == io.EOF {
					break
				}
				c.String(http.StatusInternalServerError, "Failed to read file")
				return
			}

			_, err = c.Writer.Write(buf[:n])
			if err != nil {
				c.String(http.StatusInternalServerError, "Failed to write response")
				return
			}

			c.Writer.Flush()

			select {
			case <-c.Writer.CloseNotify():
				log.Println("Client connection closed, aborting download")
				return
			default:
				continue
			}
		}
	} else {
		c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filepath.Base(filePath)))
		c.Header("Content-Length", fmt.Sprintf("%d", fileSize))
		_, err = io.Copy(c.Writer, file)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to download file")
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"msg": "file download successfully"})
}

func parseRangeHeader(rangeHeader string, fileSize int64) (rangeBytes, error) {
	rangeStr := strings.ReplaceAll(rangeHeader, "bytes=", "")
	ranges := strings.Split(rangeStr, "-")

	if len(ranges) != 2 {
		return rangeBytes{}, errors.New("invalid range header")
	}

	start, err := strconv.ParseInt(ranges[0], 10, 64)
	if err != nil {
		return rangeBytes{}, errors.New("invalid range header")
	}

	end, err := strconv.ParseInt(ranges[1], 10, 64)
	if err != nil {
		return rangeBytes{}, errors.New("invalid range header")
	}

	if start >= fileSize || end >= fileSize || start < 0 || end < 0 || start > end {
		return rangeBytes{}, errors.New("invalid range header")
	}

	length := end - start + 1
	if length > fileSize {
		length = fileSize
	}

	return rangeBytes{
		start:  start,
		end:    end,
		length: length,
	}, nil
}
