package data

import (
    "ResourceManage/model"
    "time"
    "log"
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
