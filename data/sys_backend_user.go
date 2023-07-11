package data

import (
	"ResourceManage/model"
	"gorm.io/gorm"
	"log"
)

func CreateUser(user *model.SysBackendUser, db *gorm.DB) error {
	user.Status = 1
	if err := db.Table("sys_user_backend_temp").Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUser(id string, db *gorm.DB) (*model.SysBackendUser, error) {
	var user model.SysBackendUser
	if err := db.Table("sys_user_backend_temp").Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserList(db *gorm.DB) ([]model.SysBackendUser, error) {
	var user []model.SysBackendUser
	if err := db.Table("sys_user_backend_temp").Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(id string, user *model.SysBackendUser, db *gorm.DB) error {
	existingUser, err := GetUser(id, db)
	if err != nil {
		log.Println("GetUser err:", err)
		return err
	}
	existingUser.RealName = user.RealName
	existingUser.UserPwd = user.UserPwd
	existingUser.IsSuper = user.IsSuper
	existingUser.Status = user.Status
	existingUser.Mobile = user.Mobile
	existingUser.Email = user.Email
	if err := db.Table("sys_user_backend_temp").Save(existingUser).Error; err != nil {
		return err
	}

	return nil
}

func DeleteUser(id string, db *gorm.DB) error {
	if err := db.Table("sys_user_backend_temp").Where("id = ?", id).Delete(&model.SysBackendUser{}).Error; err != nil {
		return err
	}
	return nil
}
