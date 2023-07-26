// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSysBackendUser = "sys_backend_user"

// SysBackendUser mapped from table <sys_backend_user>
type SysBackendUser struct {
	ID            int64   `gorm:"column:id;type:int;primaryKey" json:"id"`
	RealName      string  `gorm:"column:real_name;type:varchar(50);not null;comment:姓名" json:"real_name"`
	UserName      string  `gorm:"column:user_name;type:varchar(50);not null;comment:账号名" json:"user_name"`
	UserPwd       string  `gorm:"column:user_pwd;type:varchar(50);not null;comment:账号密码" json:"user_pwd"`
	IsSuper       int64   `gorm:"column:is_super;type:int;not null;comment:是否为超级管理员" json:"is_super"`
	Status        int64   `gorm:"column:status;type:int;not null;comment:状态" json:"status"`
	Mobile        string  `gorm:"column:mobile;type:varchar(16);not null;comment:手机号" json:"mobile"`
	Email         string  `gorm:"column:email;type:varchar(50);not null;comment:邮箱" json:"email"`
	Avatar        *string `gorm:"column:avatar;type:varchar(50)" json:"avatar"`
	UserType      *int64  `gorm:"column:user_type;type:int;comment:1-管理员" json:"user_type"`
	OperatorID    *int64  `gorm:"column:operator_id;type:int" json:"operator_id"`
	ParentID      *int64  `gorm:"column:parent_id;type:int;comment:上级id\n" json:"parent_id"`
	ParentRoute   *string `gorm:"column:parent_route;type:varchar(100);comment:所有上级路径" json:"parent_route"`
	Percentage    *string `gorm:"column:percentage;type:varchar(12)" json:"percentage"`
	Quota         *string `gorm:"column:quota;type:varchar(12)" json:"quota"`
	KdxfLoginName *string `gorm:"column:kdxf_login_name;type:varchar(60);comment:科大讯飞点对点登陆账号名" json:"kdxf_login_name"`
	Expires       *int64  `gorm:"column:expires;type:int;comment:账号到期时间(day)\n空为永久有效" json:"expires"`
	UnitID        *int64  `gorm:"column:unit_id;type:int" json:"unit_id"`
}

// TableName SysBackendUser's table name
func (*SysBackendUser) TableName() string {
	return TableNameSysBackendUser
}
