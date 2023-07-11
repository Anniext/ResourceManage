// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameAvtUnit = "avt_unit"

// AvtUnit mapped from table <avt_unit>
type AvtUnit struct {
	ID         int64     `gorm:"column:id;type:int;primaryKey;autoIncrement:true;comment:单位id" json:"id,string"`
	Name       string    `gorm:"column:name;type:varchar(50);not null;comment:单位名" json:"name"`
	Level      int64     `gorm:"column:level;type:int;not null;comment:单位级别" json:"level"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime;not null;comment:创建时间" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;type:datetime;not null;comment:更新时间" json:"update_time"`
	ParentID   int64     `gorm:"column:parent_id;type:int;not null;comment:上级单位id" json:"parent_id"`
	Address    string    `gorm:"column:address;type:varchar(50);not null;comment:单位地址" json:"address"`
	UserID     int64     `gorm:"column:user_id;type:int;not null;comment:用户id" json:"user_id"`
}

// TableName AvtUnit's table name
func (*AvtUnit) TableName() string {
	return TableNameAvtUnit
}
