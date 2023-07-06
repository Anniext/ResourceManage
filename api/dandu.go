package services

import (
	"ResourceManage/data"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	db := c.MustGet("db").(*gorm.DB)
	var file data.AvtFile
	//w := c.Writer
	read := c.Request
	f, handler, err := read.FormFile("file")
	if err != nil {
		log.Println("Error Retrieving the File", err)
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			log.Println("Error close file", err)
		}
	}(f)
	if err := data.UploadFile(handler, f, &file, db); err != nil {
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

	c.JSON(http.StatusOK, gin.H{"message": "file download successfully"})
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

// 批量下载
//
//type PathBody struct {
//	Path []string `json:"path"`
//}
//
//func (r *RouterGroup) Download(c *gin.Context) {
//	var filePaths PathBody
//	// 请求响应绑定File结构
//	if err := c.ShouldBind(&filePaths); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		// 错误信息400,把error发送
//		return
//	}
//	if len(filePaths.Path) == 0 {
//		c.String(http.StatusBadRequest, "File paths are missing")
//		return
//	}
//
//	var wg sync.WaitGroup
//	errChan := make(chan error, len(filePaths.Path))
//
//	for _, filePath := range filePaths.Path {
//		wg.Add(1)
//		go func(filePath string) {
//			defer wg.Done()
//
//			file, err := os.Open(filePath)
//			if err != nil {
//				errChan <- fmt.Errorf("Failed to open file %s: %v", filePath, err)
//				return
//			}
//			defer file.Close()
//
//			fileInfo, err := file.Stat()
//			if err != nil {
//				errChan <- fmt.Errorf("Failed to get file info for %s: %v", filePath, err)
//				return
//			}
//
//			fileSize := fileInfo.Size()
//
//			rangeHeader := c.Request.Header.Get("Range")
//			if rangeHeader != "" {
//				start, end, err := parseRangeHeader(rangeHeader, fileSize)
//				if err != nil {
//					errChan <- fmt.Errorf("Failed to parse Range header: %v", err)
//					return
//				}
//				fileSize = end - start + 1
//				_, err = file.Seek(start, io.SeekStart)
//				if err != nil {
//					errChan <- fmt.Errorf("Failed to seek file to start position: %v", err)
//					return
//				}
//				c.Header("Content-Range", fmt.Sprintf("bytes %d-%d/%d", start, end, fileInfo.Size()))
//				c.Writer.WriteHeader(http.StatusPartialContent)
//			} else {
//				c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filepath.Base(filePath)))
//				c.Header("Content-Type", "application/octet-stream")
//				c.Header("Content-Length", fmt.Sprintf("%d", fileSize))
//				c.Writer.WriteHeader(http.StatusOK)
//			}
//
//			_, err = io.CopyN(c.Writer, file, fileSize)
//			if err != nil {
//				errChan <- fmt.Errorf("Failed to copy file %s to response: %v", filePath, err)
//				return
//			}
//		}(filePath)
//	}
//
//	go func() {
//		wg.Wait()
//		close(errChan)
//	}()
//
//	for err := range errChan {
//		log.Println(err)
//	}
//
//	c.JSON(http.StatusOK, gin.H{"message": "files download successfully"})
//}
//
//func parseRangeHeader(rangeHeader string, fileSize int64) (int64, int64, error) {
//	rangeParts := strings.Split(rangeHeader, "=")
//	if len(rangeParts) != 2 || rangeParts[0] != "bytes" {
//		return 0, 0, fmt.Errorf("Invalid Range header format")
//	}
//
//	rangeValues := strings.Split(rangeParts[1], "-")
//	if len(rangeValues) != 2 {
//		return 0, 0, fmt.Errorf("Invalid Range header format")
//	}
//
//	start, err := strconv.ParseInt(rangeValues[0], 10, 64)
//	if err != nil || start < 0 || start >= fileSize {
//		return 0, 0, fmt.Errorf("Invalid Range header format")
//	}
//
//	end, err := strconv.ParseInt(rangeValues[1], 10, 64)
//	if err != nil || end < start || end >= fileSize {
//		return 0, 0, fmt.Errorf("Invalid Range header format")
//	}
//
//	return start, end, nil
//}
