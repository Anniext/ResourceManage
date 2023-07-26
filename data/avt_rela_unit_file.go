package data

import (
	"ResourceManage/model"
	"ResourceManage/query"
	"log"
	"strconv"
	"time"
)

type RelaList struct {
	Relas []model.RelaUnitFile
	Count int64
	Error error
}

func CreateRela(rela *model.RelaUnitFile) (bool, string) {
	rela.CreateTime = time.Now()
	rela.UpdateTime = time.Now()
	var cache *model.RelaUnitFile
	query.RelaUnitFile.Where(query.RelaUnitFile.UnitID.Eq(rela.UnitID), query.RelaUnitFile.FileID.Eq(rela.FileID)).Scan(&cache)
	if cache != nil {
		return false, "relaunitfile already exists"
	}
	if err := CacheRelaUnitFile.Sync(rela); err != nil {
		log.Println("CacheFile Sync err:", err)
		return false, err.Error()
	}
	return true, ""
}

func GetRelaList(arg *GetHeadBody, idt int64, target string) interface{} {
	var relalist RelaList
	//offset, _ := strconv.Atoi(arg.Offset)
	limit, _ := strconv.Atoi(arg.Limit)
	page, _ := strconv.Atoi(arg.Page)
	//v_delete, _ := strconv.ParseInt(arg.Delete, 10, 64)
	if target == "file" {
		if relalist.Error = query.RelaUnitFile.Select(query.RelaUnitFile.UnitID, query.RelaUnitFile.FileID).Offset((page * limit) - limit).Limit(limit).Where(query.RelaUnitFile.FileID.Eq(idt)).
			Scan(&relalist.Relas); relalist.Error != nil {
			return relalist
		}
		relalist.Count, relalist.Error = query.RelaUnitFile.Where(query.RelaUnitFile.UnitID.Eq(idt)).Count()
	} else if target == "unit" {
		if relalist.Error = query.RelaUnitFile.Select(query.RelaUnitFile.UnitID, query.RelaUnitFile.FileID).Offset((page * limit) - limit).Limit(limit).Where(query.RelaUnitFile.UnitID.Eq(idt)).
			Scan(&relalist.Relas); relalist.Error != nil {
			return relalist
		}
		relalist.Count, relalist.Error = query.RelaUnitFile.Where(query.RelaUnitFile.UnitID.Eq(idt)).Count()
	}
	if relalist.Error != nil {
		return relalist
	}
	return relalist
}

func DeleteRela(file string, unit string) (bool, string) {
	fileId, _ := strconv.ParseInt(file, 10, 64)
	unitId, _ := strconv.ParseInt(unit, 10, 64)
	var re *model.RelaUnitFile
	query.RelaUnitFile.Where(query.RelaUnitFile.FileID.Eq(fileId),
		query.RelaUnitFile.UnitID.Eq(unitId)).Scan(&re)
	if _, err := query.RelaUnitFile.Delete(re); err != nil {
		return false, err.Error()
	}
	CacheUnit.Clear()
	if err := GetUnitData(); err != nil {
		return false, "GetUnitData err:" + err.Error()
	}
	return true, ""
}

func GetRelaUnitFile(id int64, target string) (interface{}, string) {
	if target == "unit" {
		var unitList []*model.AvtUnit
		query.AvtUnit.Where(query.AvtUnit.ID.Eq(id)).Scan(&unitList)
		return unitList, ""
	} else if target == "file" {
		var fileList []*model.AvtFile
		query.AvtFile.Where(query.AvtFile.ID.Eq(id)).Scan(&fileList)
		return fileList, ""
	}
	return nil, "File does not exist"
}

func UpdateRela(unitId string, rela *model.RelaUnitFile) (bool, string) {
	uid, _ := strconv.ParseInt(unitId, 10, 64)
	var re model.RelaUnitFile
	query.RelaUnitFile.Where(query.RelaUnitFile.UnitID.Eq(uid)).Scan(&re)
	re.FileID = rela.FileID
	re.UpdateTime = time.Now()
	return true, ""
	// 省略1w字
}
