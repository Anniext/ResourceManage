package data

import (
	"gorm.io/gorm"
	"log"
	"time"
)

type AvtUnit struct {
	Id         int    `gorm:"column:id;auto" json:"id"`
	Name       string `gorm:"column:name;not null" json:"name"`
	Level      int    `gorm:"column:level" json:"level"`
	CreateTime string `gorm:"column:create_time;not null" json:"create_time"`
	UpdateTime string `gorm:"column:update_time;not null" json:"update_time"`
	IsSuper    int    `gorm:"column:is_super;not null" json:"is_super"`
	ParentId   int    `gorm:"column:parent_id;not null" json:"parent_id"`
	Address    string `gorm:"column:address;not null" json:"address"`
	//FilList    []*AvtFile
}

func LoadUnitData(db *gorm.DB) (fileDataList []*AvtUnit) {
	err := db.Raw("SELECT * FROM avt_unit ORDER BY id").Scan(&fileDataList).Error
	if err != nil {
		log.Println("avt_unit表数据加载错误：", err)
	}
	count := int64(len(fileDataList))
	if count > 0 {
		log.Println("avt_unit表缓存数据加载成功!")
		for _, unit := range fileDataList {
			CacheUnit.Set(unit)
		}
	} else {
		log.Println("没有数据！")
	}
	return
}

func (m *UnitMap) Set(bu *AvtUnit) {
	m.lock.Lock()
	m.data[bu.Id] = bu
	m.lock.Unlock()
}

func CreateUnit(unit *AvtUnit, db *gorm.DB) error {
	unit.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	unit.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
	if err := db.Table("avt_unit").Create(unit).Error; err != nil {
		return err
	}
	return nil
}

func GetUnit(id string, db *gorm.DB) (*AvtUnit, error) {
	var unit AvtUnit
	if err := db.Table("avt_unit").Where("id = ?", id).First(&unit).Error; err != nil {
		return nil, err
	}
	return &unit, nil
}

func GetUnitList(db *gorm.DB, delete_id string) ([]AvtUnit, error) {
	var units []AvtUnit
	if delete_id == "0" {
		if err := db.Table("avt_unit").Find(&units).Error; err != nil {
			return nil, err
		}
	} else if delete_id == "1" {
		if err := db.Table("avt_unit").Where("is_delete = ?", 1).Find(&units).Error; err != nil {
			return nil, err
		}
	} else if delete_id == "2" {
		if err := db.Table("avt_unit").Where("is_delete = ?", 0).Find(&units).Error; err != nil {
			return nil, err
		}
	}
	return units, nil
}

func UpdateUnit(id string, unit *AvtUnit, db *gorm.DB) error {
	existingUnit, err := GetUnit(id, db)
	if err != nil {
		log.Println("GetUnit err:", err)
		return err
	}
	existingUnit.Name = unit.Name
	existingUnit.Level = unit.Level
	existingUnit.IsSuper = unit.IsSuper
	existingUnit.ParentId = unit.ParentId
	existingUnit.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
	existingUnit.Address = unit.Address
	if err := db.Table("avt_unit").Save(existingUnit).Error; err != nil {
		return err
	}

	return nil
}

func DeleteUnit(id string, db *gorm.DB) error {
	if err := db.Table("avt_unit").Where("id = ?", id).Delete(&AvtUnit{}).Error; err != nil {
		return err
	}
	return nil
}
