package data

import (
	"gorm.io/gorm"
	"log"
)

type SysBackendUser struct {
	Id            int    `gorm:"column:id;auto" json:"id"`
	RealName      string `gorm:"column:real_name;not null" json:"real_name"`
	UserName      string `gorm:"column:user_name;not null" json:"user_name"`
	UserPwd       string `gorm:"column:user_pwd;not null" json:"user_pwd"`
	IsSuper       int    `gorm:"column:is_super" json:"is_super"`
	Status        int    `gorm:"column:status;not null" json:"status"`
	Mobile        string `gorm:"column:mobile" json:"mobile"`
	Email         string `gorm:"column:email" json:"email"`
	Avatar        string `gorm:"column:avatar" json:"avatar"`
	UserType      int    `gorm:"column:user_type;not null" json:"user_type"`
	OperatorId    int    `gorm:"column:operator_id" json:"operator_id"`
	ParentId      int    `gorm:"column:parent_id" json:"parent_id"`
	ParentRoute   string `gorm:"column:parent_route" json:"parent_route"`
	Percentage    int    `gorm:"column:percentage" json:"percentage"`
	Quota         int    `gorm:"column:quota" json:"quota"`
	KdxfLoginName string `gorm:"column:kdxf_login_name" json:"kdxf_login_name"`
	Expires       int    `gorm:"column:expires" json:"expires"`
}

func LoadBackendUserData(db *gorm.DB) (UserDataList []*SysBackendUser) {
	err := db.Raw("SELECT * FROM sys_backend_user ORDER BY id").Scan(&UserDataList).Error
	if err != nil {
		log.Println("sys_backend_user表数据加载错误：", err)
	}
	count := int64(len(UserDataList))
	if count > 0 {
		log.Println("sys_backend_user表缓存数据加载成功!")
		for _, user := range UserDataList {
			CacheBackendUser.Set(user)
		}
	} else {
		log.Println("没有数据！")
	}
	return
}

func (m *BackendUserMap) Set(bu *SysBackendUser) {
	m.lock.Lock()
	m.data[bu.UserName] = bu
	m.lock.Unlock()
}

func (m *BackendUserMap) Get(name string) *SysBackendUser {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.data[name]
}
