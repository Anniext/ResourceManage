package data

type SysBackendUser struct {
	Id            int    `gorm:"column:id;auto" json:"id"`
	RealName      string `gorm:"column:real_name;not null" json:"real_name"`
	UserName      string `gorm:"column:user_name;not null" json:"user_name"`
	UserPwd       string `gorm:"column:user_pwd;not null" json:"user_pwd"`
	IsSuper       int    `gorm:"column:is_super;not null" json:"is_super"`
	Status        int    `gorm:"column:status;not null" json:"status"`
	Mobile        string `gorm:"column:mobile;not null" json:"mobile"`
	Email         string `gorm:"column:email;not null" json:"email"`
	Avatar        string `gorm:"column:avatar;not null" json:"avatar"`
	UserYype      int    `gorm:"column:user_type;not null" json:"user_type"`
	OperatorId    int    `gorm:"column:operator_id;not null" json:"operator_id"`
	ParentId      int    `gorm:"column:parent_id;not null" json:"parent_id"`
	ParentRoute   string `gorm:"column:parent_route;not null" json:"parent_route"`
	Percentage    int    `gorm:"column:percentage;not null" json:"percentage"`
	Quota         int    `gorm:"column:quota;not null" json:"quota"`
	KdxfLoginName string `gorm:"column:kdxf_login_name;not null" json:"kdxf_login_name"`
}
