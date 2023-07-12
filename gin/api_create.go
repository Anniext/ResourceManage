package services

import (
	"ResourceManage/data"
	"ResourceManage/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func (r *RouterGroup) Create(c *gin.Context) {
	groupName := GetGroupName(r)
	switch groupName {
	case "resource":
		FileCreateGroup(c)
	case "unit":
		UnitCreateGroup(c)
	case "user":
		UserCreateGroup(c)
	default:
		log.Println("Error group name", groupName)
	}
}

// FileCreateGroup 创建资源
func FileCreateGroup(c *gin.Context) {
	var file model.AvtFile
	// 请求响应绑定File结构
	if err := c.ShouldBind(&file); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 错误信息400,把error发送
		return
	}

	if err := data.CreateFile(&file); err != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	// 发送状态码200
	c.JSON(http.StatusOK, gin.H{"msg": "File created successfully"})

}

func UnitCreateGroup(c *gin.Context) {
	//level, errStr := utils.GetLevel(utils.GetJwtClaims(c)) //通过token获取level
	//if errStr != "" {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": errStr})
	//	return
	//}
	var unit model.AvtUnit
	// 请求响应绑定File结构
	if err := c.ShouldBind(&unit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 错误信息400,把error发送
		return
	}
	//if unit.Level <= level {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": "Permission too low to create"})
	//	return
	//}
	if err := data.CreateUnit(&unit); err != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		// 错误信息500,把error发送
		return
	}
	// 发送状态码200
	c.JSON(http.StatusOK, gin.H{"msg": "Unit created successfully"})
}

func UserCreateGroup(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var user model.SysBackendUser
	// 请求响应绑定File结构
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 错误信息400,把error发送
		return
	}
	if err := data.CreateUser(&user, db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		// 错误信息500,把error发送
		return
	}
	// 发送状态码200
	c.JSON(http.StatusOK, "User created successfully")
}
