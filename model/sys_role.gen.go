// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSysRole = "sys_role"

// SysRole mapped from table <sys_role>
type SysRole struct {
	ID       int64  `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id,string"`
	Name     string `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Seq      int64  `gorm:"column:seq;type:int;not null" json:"seq"`
	RoleType *int64 `gorm:"column:role_type;type:int;comment:0-系统管理员角色；1-企业客户角色；2-运营商角色" json:"role_type"`
}

// TableName SysRole's table name
func (*SysRole) TableName() string {
	return TableNameSysRole
}