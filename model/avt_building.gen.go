// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameAvtBuilding = "avt_building"

// AvtBuilding mapped from table <avt_building>
type AvtBuilding struct {
	ID       int64   `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id,string"`
	Name     *string `gorm:"column:name;type:varchar(30)" json:"name"`
	IsDelete *int64  `gorm:"column:is_delete;type:int;comment:0-未删除；1-删除" json:"is_delete"`
	Sort     *int64  `gorm:"column:sort;type:int;comment:越大越靠前" json:"sort"`
}

// TableName AvtBuilding's table name
func (*AvtBuilding) TableName() string {
	return TableNameAvtBuilding
}