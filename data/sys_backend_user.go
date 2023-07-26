package data

import (
	"ResourceManage/model"
	"ResourceManage/query"
	"gorm.io/gorm"
	"strconv"
)

func CreateUser(user *model.SysBackendUser, db *gorm.DB) error {
	user.Status = 1
	if err := db.Table("sys_user_backend_temp").Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUser(user *model.SysBackendUser, name string) (bool, string) {
	//if user := CacheBackendUser.Get(name); user != nil {
	//	return true, ""
	//}
	if err := query.SysBackendUser.Where(query.SysBackendUser.UserName.Eq(name)).Scan(user); err != nil {
		return false, "File does not exist"
	}
	return true, ""
}

func GetUserList(id string) ([]model.SysBackendUser, error) {
	var user []model.SysBackendUser
	if id == "" {
		if err := query.SysBackendUser.Where(query.SysBackendUser.UnitID.IsNull()).Scan(&user); err != nil {
			return nil, err
		}
		return user, nil
	}
	idt, _ := strconv.ParseInt(id, 10, 64)
	if err := query.SysBackendUser.Where(query.SysBackendUser.UnitID.Eq(idt)).Scan(&user); err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(name string, user *model.SysBackendUser) string {
	var existingUser model.SysBackendUser
	ok, err := GetUser(&existingUser, name)
	if !ok {
		return err
	}
	existingUser.UnitID = user.UnitID
	CacheBackendUser.Update(&existingUser)
	return ""

}

func DeleteUser(id string, db *gorm.DB) error {
	if err := db.Table("sys_user_backend_temp").Where("id = ?", id).Delete(&model.SysBackendUser{}).Error; err != nil {
		return err
	}
	return nil
}
