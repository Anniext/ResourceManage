// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameAvtCourseChapter = "avt_course_chapter"

// AvtCourseChapter mapped from table <avt_course_chapter>
type AvtCourseChapter struct {
	ID            int64      `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id,string"`
	CourseID      int64      `gorm:"column:course_id;type:int;not null;comment:课程id" json:"course_id"`
	Name          string     `gorm:"column:name;type:varchar(300);not null;comment:节点名称" json:"name"`
	ParentID      int64      `gorm:"column:parentId;type:int;not null;comment:父级ID" json:"parentId"`
	Type          int64      `gorm:"column:type;type:int;not null;comment:0课时节点1是章节" json:"type"`
	Status        int64      `gorm:"column:status;type:int;not null;comment:状态0正常1删除" json:"status"`
	CreateTime    time.Time  `gorm:"column:create_time;type:datetime;not null;comment:添加时间" json:"create_time"`
	Sort          int64      `gorm:"column:sort;type:int;not null;comment:显示排序" json:"sort"`
	PlayCount     int64      `gorm:"column:play_count;type:int;not null;comment:播放次数" json:"play_count"`
	Isfree        int64      `gorm:"column:isfree;type:int;not null;comment:是否可以试听1免费2收费" json:"isfree"`
	Videotype     *string    `gorm:"column:videotype;type:varchar(20)" json:"videotype"`
	Videourl      string     `gorm:"column:videourl;type:varchar(500);not null;comment:视频地址" json:"videourl"`
	Videojson     *string    `gorm:"column:videojson;type:text;comment:json格式辅助56" json:"videojson"`
	TeacherID     int64      `gorm:"column:teacher_id;type:int;not null;comment:讲师id" json:"teacher_id"`
	CourseHour    *int64     `gorm:"column:course_hour;type:int" json:"course_hour"`
	CourseMinutes int64      `gorm:"column:course_minutes;type:int;not null;comment:时长：分钟" json:"course_minutes"`
	CourseSeconds int64      `gorm:"column:course_seconds;type:int;not null;comment:时长：秒" json:"course_seconds"`
	FileType      *string    `gorm:"column:file_type;type:varchar(20);comment:VIDEO视频 AUDIO音频 FILE文档 TXT文本 ATLAS图片集" json:"file_type"`
	Courseware    *string    `gorm:"column:courseware;type:varchar(255);comment:讲义下载地址" json:"courseware"`
	Content       *string    `gorm:"column:content;type:text;comment:文本" json:"content"`
	PageCount     *int64     `gorm:"column:page_count;type:int;comment:页数" json:"page_count"`
	ExamLink      *string    `gorm:"column:exam_link;type:varchar(255);comment:è€ƒè¯•è¯•å·é“¾æŽ¥" json:"exam_link"`
	TouristIsfree *int64     `gorm:"column:tourist_isfree;type:int;comment:æ¸¸å®¢è¯•çœ‹ 1æ˜¯ 2å¦" json:"tourist_isfree"`
	VideoPwd      *string    `gorm:"column:video_pwd;type:varchar(255);comment:点播密码" json:"video_pwd"`
	StartTime     *int64     `gorm:"column:startTime;type:bigint;comment:直播开始" json:"startTime"`
	EndTime       *int64     `gorm:"column:endTime;type:bigint;comment:直播结束" json:"endTime"`
	UpdateTime    *time.Time `gorm:"column:update_time;type:timestamp;default:CURRENT_TIMESTAMP" json:"update_time"`
	Nbtype        *string    `gorm:"column:nbtype;type:varchar(50);comment:牛邦视频类型" json:"nbtype"`
	IsDisplay     *int64     `gorm:"column:is_display;type:int;default:1;comment:章节是否显示在课程目录中" json:"is_display"`
	Logo          *string    `gorm:"column:logo;type:varchar(200);comment:图片路径" json:"logo"`
	ClassID       *string    `gorm:"column:class_id;type:varchar(40)" json:"class_id"`
	RoomID        *int64     `gorm:"column:room_id;type:int;comment:对于直播课，需要填写上课教室" json:"room_id"`
}

// TableName AvtCourseChapter's table name
func (*AvtCourseChapter) TableName() string {
	return TableNameAvtCourseChapter
}