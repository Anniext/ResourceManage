// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameAvtDeviceState = "avt_device_states"

// AvtDeviceState mapped from table <avt_device_states>
type AvtDeviceState struct {
	ID         int64      `gorm:"column:id;type:int;primaryKey" json:"id,string"`
	RoomID     int64      `gorm:"column:room_id;type:int;not null" json:"room_id"`
	CreateTime *time.Time `gorm:"column:create_time;type:datetime;comment:上课时间" json:"create_time"`
	UpdateTime *time.Time `gorm:"column:update_time;type:datetime;comment:下课时间" json:"update_time"`
	Interval   *int64     `gorm:"column:interval;type:int;comment:时间间隔: 更新时间-创建时间" json:"interval"`
	JoinNum    *int64     `gorm:"column:join_num;type:int" json:"join_num"`
}

// TableName AvtDeviceState's table name
func (*AvtDeviceState) TableName() string {
	return TableNameAvtDeviceState
}
