// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameAvtPayCardLog = "avt_pay_card_log"

// AvtPayCardLog mapped from table <avt_pay_card_log>
type AvtPayCardLog struct {
	ID         int64   `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id,string"`
	CardNo     *string `gorm:"column:card_no;type:varchar(30)" json:"card_no"`
	ActionTime *string `gorm:"column:action_time;type:varchar(20)" json:"action_time"`
	Result     *int64  `gorm:"column:result;type:int;comment:0刷卡成功  1无效卡  2时间错误  3其他错误" json:"result"`
	RoomID     *int64  `gorm:"column:room_id;type:int" json:"room_id"`
}

// TableName AvtPayCardLog's table name
func (*AvtPayCardLog) TableName() string {
	return TableNameAvtPayCardLog
}
