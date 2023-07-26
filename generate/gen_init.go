package generate

import (
	"ResourceManage/config"
	"ResourceManage/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
	"strings"
)

type RelateConfig struct {
	// specify field's type
	RelatePointer      bool // ex: CreditCard  *CreditCard
	RelateSlice        bool // ex: CreditCards []CreditCard
	RelateSlicePointer bool // ex: CreditCards []*CreditCard

	JSONTag      string // related field's JSON tag
	GORMTag      string // related field's GORM tag
	NewTag       string // related field's new tag
	OverwriteTag string // related field's tag
}

var MySQLDSN = config.Configs.Dev.DSN

func GenGenerate() {
	db, err := gorm.Open(mysql.Open(MySQLDSN))
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}

	// 构造生成器实例
	g := gen.NewGenerator(gen.Config{
		OutPath:           "./query",
		ModelPkgPath:      "./model",
		Mode:              gen.WithDefaultQuery | gen.WithoutContext | gen.WithQueryInterface,
		FieldNullable:     true,
		FieldCoverable:    false,
		FieldSignable:     false,
		FieldWithIndexTag: false,
		FieldWithTypeTag:  true,
	})
	g.UseDB(db)
	dataMap := map[string]func(columnType gorm.ColumnType) (dataType string){
		"tinyint":   func(columnType gorm.ColumnType) (dataType string) { return "int64" },
		"smallint":  func(columnType gorm.ColumnType) (dataType string) { return "int64" },
		"mediumint": func(columnType gorm.ColumnType) (dataType string) { return "int64" },
		"bigint":    func(columnType gorm.ColumnType) (dataType string) { return "int64" },
		"int":       func(columnType gorm.ColumnType) (dataType string) { return "int64" },
	}

	g.WithDataTypeMap(dataMap)
	jsonField := gen.FieldJSONTagWithNS(func(columnName string) (tagContent string) {
		toStringField := `id, `
		if strings.Contains(toStringField, columnName) {
			return columnName + ",string"
		}
		return columnName
	})

	//autoUpdateTimeField := gen.FieldGORMTag("update_time", "column:update_time;type:int unsigned;autoUpdateTime")
	//autoCreateTimeField := gen.FieldGORMTag("create_time", "column:create_time;type:int unsigned;autoCreateTime")
	softDeleteField := gen.FieldType("delete_time", "gorm.DeletedAt")
	// 模型自定义选项组
	//fieldOpts := []gen.ModelOpt{jsonField, autoCreateTimeField, autoUpdateTimeField, softDeleteField}
	fieldOpts := []gen.ModelOpt{jsonField, softDeleteField}

	//Unit := g.GenerateModel("avt_unit",
	//	append(
	//		fieldOpts,
	//		// user 一对多 address 关联, 外键`uid`在 address 表中
	//		gen.FieldRelate(field.HasMany, "Unit", Unit, &field.RelateConfig{GORMTag: keytag}),
	//	)...)
	//allModel := g.GenerateAllTable(fieldOpts...)
	//g.ApplyBasic(allModel...)

	keytag1 := make(map[string]string)
	keytag2 := make(map[string]string)
	keytag3 := make(map[string]string)
	avt_user := g.GenerateModel("sys_backend_user")
	rela := g.GenerateModel("rela_unit_file")

	keytag1["foreignKey"] = "UnitID"
	keytag2["foreignKey"] = "ParentID"
	avt_unit := g.GenerateModel("avt_unit",
		append(
			fieldOpts,
			gen.FieldRelate(field.HasMany, "FileList", rela,
				&field.RelateConfig{
					GORMTag: keytag1,
				}),
			gen.FieldRelate(field.HasMany, "UserList", avt_user,
				&field.RelateConfig{GORMTag: keytag1}),
			gen.FieldRelateModel(field.HasMany, "SubUnitList", model.AvtUnit{},
				&field.RelateConfig{GORMTag: keytag2}),
		)...)

	keytag3["foreignKey"] = "FileID"
	avt_file := g.GenerateModel("avt_file",
		append(
			fieldOpts,
			gen.FieldRelate(field.HasMany, "UnitList", rela,
				&field.RelateConfig{
					GORMTag: keytag3,
				}),
		)...)

	g.ApplyBasic(avt_user, avt_unit, rela, avt_file)

	g.Execute()
}
