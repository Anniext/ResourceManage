package data

import (
	"ResourceManage/model"
	"gorm.io/gorm"
)

func CreateUser(user *model.SysBackendUser, db *gorm.DB) error {
	user.Status = 1
	if err := db.Table("sys_user_backend_temp").Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUser(name string) (*model.SysBackendUser, string) {
	if user := CacheBackendUser.Get(name); user != nil {
		return user, ""
	}
	return nil, "File does not exist"
}

func GetUserList(db *gorm.DB) ([]model.SysBackendUser, error) {
	var user []model.SysBackendUser
	if err := db.Table("sys_user_backend_temp").Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(name string, user *model.SysBackendUser, db *gorm.DB) string {
	existingUser, err := GetUser(name)
	if err != "" {
		return err
	}
	existingUser.UserName = user.UserName
	existingUser.Status = user.Status
	existingUser.Email = user.Email
	existingUser.Level = user.Level
	CacheBackendUser.Update(existingUser)
	return ""

}

func DeleteUser(id string, db *gorm.DB) error {
	if err := db.Table("sys_user_backend_temp").Where("id = ?", id).Delete(&model.SysBackendUser{}).Error; err != nil {
		return err
	}
	return nil
}
