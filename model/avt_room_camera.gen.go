// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameAvtRoomCamera = "avt_room_camera"

// AvtRoomCamera mapped from table <avt_room_camera>
type AvtRoomCamera struct {
	ID            int64   `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id,string"`
	Title         *string `gorm:"column:title;type:varchar(50)" json:"title"`
	RoomID        *int64  `gorm:"column:room_id;type:int" json:"room_id"`
	RtspURL       *string `gorm:"column:rtsp_url;type:varchar(255)" json:"rtsp_url"`
	Group_        *int64  `gorm:"column:group;type:int;comment:播放窗口分组" json:"group"`
	FlvURL        *string `gorm:"column:flv_url;type:varchar(255);comment:web播放的flv地址" json:"flv_url"`
	DirectionType *int64  `gorm:"column:direction_type;type:int;comment:0-pull;1-push" json:"direction_type"`
	Enable        *int64  `gorm:"column:enable;type:int;default:1;comment:1-启用；0-禁用" json:"enable"`
	Type          *int64  `gorm:"column:type;type:int;comment:1-老师通道；2-学生通道；3-讲台电脑" json:"type"`
	GbNamespace   *string `gorm:"column:gb_namespace;type:varchar(60);comment:国标空间id" json:"gb_namespace"`
	GbChannelID   *string `gorm:"column:gb_channel_id;type:varchar(60);comment:国标通道id" json:"gb_channel_id"`
}

// TableName AvtRoomCamera's table name
func (*AvtRoomCamera) TableName() string {
	return TableNameAvtRoomCamera
}