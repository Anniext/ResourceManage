// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSysRoleBackenduserRel = "sys_role_backenduser_rel"

// SysRoleBackenduserRel mapped from table <sys_role_backenduser_rel>
type SysRoleBackenduserRel struct {
	ID            int64     `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id,string"`
	RoleID        int64     `gorm:"column:role_id;type:int;not null" json:"role_id"`
	BackendUserID int64     `gorm:"column:backend_user_id;type:int;not null" json:"backend_user_id"`
	Created       time.Time `gorm:"column:created;type:datetime;not null" json:"created"`
}

// TableName SysRoleBackenduserRel's table name
func (*SysRoleBackenduserRel) TableName() string {
	return TableNameSysRoleBackenduserRel
}