package data

import (
	"ResourceManage/model"
	"ResourceManage/query"
	"log"
	"strconv"
	"time"
)

type NamePrmiss struct {
	Name   string
	Primss map[string]interface{}
}

func CreateUnit(unit *model.AvtUnit) string {
	unit.CreateTime = time.Now()
	unit.UpdateTime = time.Now()
	if CacheUnit.Get(unit.Name) != nil {
		return "Unit name already exists"
	}
	if err := CacheUnit.Sync(unit); err != nil {
		log.Println("CacheFile Sync err:", err)
		return err.Error()
	}
	CacheUnit.Set(unit)
	return ""
}

func GetUnit(name string) (*model.AvtUnit, string) {
	if unit := CacheUnit.Get(name); unit != nil {
		return unit, ""
	}
	return nil, "Unit does not exist"
}

func GetUnitList(arg *GetHeadBody, prmiss map[string]interface{}) ([]model.AvtUnit, int64, error) {
	var units []model.AvtUnit
	//offset, _ := strconv.Atoi(arg.Offset)
	limit, _ := strconv.Atoi(arg.Limit)
	page, _ := strconv.Atoi(arg.Page)
	level := int64(prmiss["level"].(float64))
	userid := int64(prmiss["user_id"].(float64))
	//v_delete, _ := strconv.ParseInt(arg.Delete, 10, 64)
	if err := query.AvtUnit.Offset((page * limit) - limit).Limit(limit).Where(query.AvtUnit.UserID.
		Eq(userid)).Where(query.AvtUnit.Level.Gt(level)).Or(query.AvtUnit.Level.Eq(level)).Scan(&units); err != nil {
		return nil, 0, err
	}
	count, err := query.AvtUnit.Where(query.AvtUnit.UserID.
		Eq(userid)).Where(query.AvtUnit.Level.Gt(level)).Or(query.AvtUnit.Level.Eq(level)).Count()
	if err != nil {
		return nil, 0, err
	}
	return units, count, nil
}

func UpdateUnit(name string, unit *model.AvtUnit) string {
	existingUnit, err := GetUnit(name)
	if err != "" {
		return err
	}
	existingUnit.Name = unit.Name
	existingUnit.UpdateTime = time.Now()
	existingUnit.Address = unit.Address
	CacheUnit.Update(existingUnit)

	return ""
}

func DeleteUnit(i interface{}) string {
	if p, ok := i.(NamePrmiss); ok {
		cache := CacheUnit.Get(p.Name)
		level := p.Primss["level"].(float64)
		if cache == nil {
			return "Unit does not exist"
		}
		if cache.Level <= int64(level) {
			return "Permission too low to delete"
		}
		if _, err := query.AvtUnit.Delete(cache); err != nil {
			log.Println("avt_unit表数据同步错误：", err)
			return err.Error()
		}
		CacheUnit.Clear()
		if err := GetUnitData(); err != nil {
			return "GetUnitData err:" + err.Error()
		}
	} else {
		return "interface Reflection err"
	}
	return ""
}
