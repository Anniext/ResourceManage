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

type UnitList struct {
    Units []model.AvtUnit
    Count int64
    Error error
}

func CreateUnit(unit *model.AvtUnit) (bool, string) {
	unit.CreateTime = time.Now()
	unit.UpdateTime = time.Now()
    id := CacheUnit.GetID(unit.Name)
	if CacheUnit.Get(id) != nil {
		return false, "Unit name already exists"
	}
	if err := CacheUnit.Sync(unit); err != nil {
		log.Println("CacheFile Sync err:", err)
		return false, err.Error()
	}
	CacheUnit.Set(unit)
	return true, ""
}

func GetUnit(id int64) (*model.AvtUnit, string) {
	if unit := CacheUnit.Get(id); unit != nil {
		return unit, ""
	}
	return nil, "Unit does not exist"
}

func GetUnitList(arg *GetHeadBody, prmiss map[string]interface{}) (interface{}) {
	var unitlist UnitList
	//offset, _ := strconv.Atoi(arg.Offset)
	limit, _ := strconv.Atoi(arg.Limit)
	page, _ := strconv.Atoi(arg.Page)
	level := int64(prmiss["level"].(float64))
	userid := int64(prmiss["user_id"].(float64))
	//v_delete, _ := strconv.ParseInt(arg.Delete, 10, 64)
	if unitlist.Error = query.AvtUnit.Offset((page * limit) - limit).Limit(limit).Where(query.AvtUnit.UserID.
		Eq(userid)).Where(query.AvtUnit.Level.Gt(level)).Or(query.AvtUnit.Level.Eq(level)).Scan(&unitlist.Units); unitlist.Error != nil {
		return unitlist
	}
	unitlist.Count, unitlist.Error = query.AvtUnit.Where(query.AvtUnit.UserID.
		Eq(userid)).Where(query.AvtUnit.Level.Gt(level)).Or(query.AvtUnit.Level.Eq(level)).Count()
	if unitlist.Error != nil {
		return unitlist
	}
	return unitlist
}

func UpdateUnit(name string, unit *model.AvtUnit) (bool, string) {
	existingUnit, err := GetUnit(CacheUnit.GetID(name))
	if err != "" {
		return false, err
	}
	existingUnit.Name = unit.Name
	existingUnit.UpdateTime = time.Now()
	existingUnit.Address = unit.Address
	CacheUnit.Update(existingUnit)
	return true, "" 
}

func DeleteUnit(i interface{}) (bool, string) {
	if p, ok := i.(NamePrmiss); ok {
		cache := CacheUnit.Get(CacheUnit.GetID(p.Name))
		level := p.Primss["level"].(float64)
		if cache == nil {
			return false, "Unit does not exist"
		}
		if cache.Level <= int64(level) {
			return false, "Permission too low to delete"
		}
		if _, err := query.AvtUnit.Delete(cache); err != nil {
			return false, err.Error()
		}
		CacheUnit.Clear()
		if err := GetUnitData(); err != nil {
			return false, "GetUnitData err:" + err.Error()
		}
	} else {
		return false, "interface Reflection err"
	}
	return true, ""
}
