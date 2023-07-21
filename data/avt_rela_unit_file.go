package data

import (
	"ResourceManage/model"
	"ResourceManage/query"
	"log"
	"strconv"
	"time"
)

func CreateRela(rela *model.RelaUnitFile) string {
	rela.CreateTime = time.Now()
	rela.UpdateTime = time.Now()
    cache := CacheRelaUnitFile.Get(rela.UnitID)
    if cache != nil && cache.FileID == rela.FileID{
        return "relaunitfile already exists"
    }
	if err := CacheRelaUnitFile.Sync(rela); err != nil {
		log.Println("CacheFile Sync err:", err)
		return err.Error()
	}
	CacheRelaUnitFile.Set(rela)
	return ""
}

func GetRelaList(arg *GetHeadBody) ([]model.RelaUnitFile, int64, error) {
	var rela []model.RelaUnitFile
	//offset, _ := strconv.Atoi(arg.Offset)
	limit, _ := strconv.Atoi(arg.Limit)
	page, _ := strconv.Atoi(arg.Page)
	//v_delete, _ := strconv.ParseInt(arg.Delete, 10, 64)
	if err := query.RelaUnitFile.Offset((page * limit) - limit).Limit(limit).Scan(&rela); err != nil {
		return nil, 0, err
	}
	count, err := query.RelaUnitFile.Count()
	if err != nil {
		return nil, 0, err
	}
	return rela, count, nil
}
