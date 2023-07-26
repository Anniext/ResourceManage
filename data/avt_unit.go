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
	Units []*model.AvtUnit
	Count int64
	Error error
}

func CreateUnit(unit *model.AvtUnit) (bool, string) {
	unit.CreateTime = time.Now()
	unit.UpdateTime = time.Now()
	if CacheUnit.Get(CacheUnit.GetID(unit.Name)) != nil {
		return false, "Unit name already exists"
	}
	if err := CacheUnit.Sync(unit); err != nil {
		log.Println("CacheFile Sync err:", err)
		return false, err.Error()
	}
	newUnit, _ := query.AvtUnit.
		Where(query.AvtUnit.Name.
			Eq(unit.Name)).
		Preload(query.AvtUnit.FileList,
			query.AvtUnit.SubUnitList, query.AvtUnit.UserList).
		First()
	CacheUnit.Set(newUnit)
	return true, ""
}

func GetUnit(id int64) (*model.AvtUnit, string) {
	if unit := CacheUnit.Get(id); unit != nil {
		return unit, ""
	}
	return nil, "Unit does not exist"
}

func GetUnitList(arg *GetHeadBody, prmiss map[string]interface{}) interface{} {
	var unitlist UnitList
	//offset, _ := strconv.Atoi(arg.Offset)
	limit, _ := strconv.Atoi(arg.Limit)
	page, _ := strconv.Atoi(arg.Page)
	unitid := int64(prmiss["unit_id"].(float64))
	var level int64
	unitlist.Error = query.AvtUnit.
		Where(query.AvtUnit.ID.
			Eq(unitid)).
		Select(query.AvtUnit.Level).
		Scan(&level)
	if unitlist.Error != nil {
		return unitlist
	}

	unitlist.Units, unitlist.Error = query.AvtUnit.
		Offset((page*limit)-limit).
		Limit(limit).
		Preload(query.AvtUnit.FileList,
			query.AvtUnit.SubUnitList, query.AvtUnit.UserList).Find()
	if unitlist.Error != nil {
		return unitlist
	}

	unitlist.Count, unitlist.Error = query.AvtUnit.
		Preload(query.AvtUnit.FileList, query.AvtUnit.SubUnitList, query.AvtUnit.UserList).
		Count()
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
		unid := int64(p.Primss["unit_id"].(float64))
		ucache, _ := query.AvtUnit.Where(query.AvtUnit.ID.Eq(unid)).First()
		if cache == nil {
			return false, "Unit does not exist"
		}
		if cache.Level <= ucache.Level {
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
