package data

import (
	"ResourceManage/model"
	"gorm.io/gorm"
	"log"
	"time"
)

func CreateUnit(unit *model.AvtUnit, db *gorm.DB) error {
	unit.CreateTime = time.Now()
	unit.UpdateTime = time.Now()
	if err := db.Table("avt_unit").Create(unit).Error; err != nil {
		return err
	}
	return nil
}

func GetUnit(id string, db *gorm.DB) (*model.AvtUnit, error) {
	var unit model.AvtUnit
	if err := db.Table("avt_unit").Where("id = ?", id).First(&unit).Error; err != nil {
		return nil, err
	}
	return &unit, nil
}

func GetUnitList(db *gorm.DB, delete_id string, l int64) ([]model.AvtUnit, error) {
	var units []model.AvtUnit
	if delete_id == "0" {
		if err := db.Table("avt_unit").Where("level >= ?", l).Find(&units).Error; err != nil {
			return nil, err
		}
	} else if delete_id == "1" {
		if err := db.Table("avt_unit").Where("is_delete = ? and level >= ?", 1, l).Find(&units).Error; err != nil {
			return nil, err
		}
	} else if delete_id == "2" {
		if err := db.Table("avt_unit").Where("is_delete = ? and level >= ?", 0, l).Find(&units).Error; err != nil {
			return nil, err
		}
	}
	return units, nil
}

func UpdateUnit(id string, unit *model.AvtUnit, db *gorm.DB) error {
	existingUnit, err := GetUnit(id, db)
	if err != nil {
		log.Println("GetUnit err:", err)
		return err
	}
	existingUnit.Name = unit.Name
	existingUnit.Level = unit.Level
	existingUnit.UserID = unit.UserID
	existingUnit.ParentID = unit.ParentID
	existingUnit.UpdateTime = time.Now()
	existingUnit.Address = unit.Address
	if err := db.Table("avt_unit").Save(existingUnit).Error; err != nil {
		return err
	}

	return nil
}

func DeleteUnit(id string, db *gorm.DB) error {
	if err := db.Table("avt_unit").Where("id = ?", id).Delete(&model.AvtUnit{}).Error; err != nil {
		return err
	}
	return nil
}
