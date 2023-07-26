// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"ResourceManage/model"
)

func newAvtEquipment(db *gorm.DB, opts ...gen.DOOption) avtEquipment {
	_avtEquipment := avtEquipment{}

	_avtEquipment.avtEquipmentDo.UseDB(db, opts...)
	_avtEquipment.avtEquipmentDo.UseModel(&model.AvtEquipment{})

	tableName := _avtEquipment.avtEquipmentDo.TableName()
	_avtEquipment.ALL = field.NewAsterisk(tableName)
	_avtEquipment.ID = field.NewInt64(tableName, "id")
	_avtEquipment.Type = field.NewInt64(tableName, "type")
	_avtEquipment.Name = field.NewString(tableName, "name")
	_avtEquipment.Brand = field.NewString(tableName, "brand")
	_avtEquipment.PurchasingDate = field.NewString(tableName, "purchasing_date")
	_avtEquipment.Model = field.NewString(tableName, "model")
	_avtEquipment.ServiceLife = field.NewInt64(tableName, "service_life")
	_avtEquipment.WorkingTime = field.NewInt64(tableName, "working_time")
	_avtEquipment.IsDelete = field.NewInt64(tableName, "is_delete")
	_avtEquipment.CreateTime = field.NewTime(tableName, "create_time")
	_avtEquipment.UpdateTime = field.NewTime(tableName, "update_time")
	_avtEquipment.BuildingID = field.NewInt64(tableName, "building_id")
	_avtEquipment.FloorID = field.NewInt64(tableName, "floor_id")
	_avtEquipment.RoomID = field.NewInt64(tableName, "room_id")
	_avtEquipment.Group_ = field.NewInt64(tableName, "group")
	_avtEquipment.IP = field.NewString(tableName, "ip")
	_avtEquipment.Status = field.NewInt64(tableName, "status")

	_avtEquipment.fillFieldMap()

	return _avtEquipment
}

type avtEquipment struct {
	avtEquipmentDo

	ALL            field.Asterisk
	ID             field.Int64
	Type           field.Int64 // 设备类型
	Name           field.String
	Brand          field.String // 品牌
	PurchasingDate field.String // 购买时间
	Model          field.String // 型号
	ServiceLife    field.Int64  // 月
	WorkingTime    field.Int64  // 使用时长
	IsDelete       field.Int64  // 0-未删除；1-删除
	CreateTime     field.Time
	UpdateTime     field.Time
	BuildingID     field.Int64
	FloorID        field.Int64
	RoomID         field.Int64
	Group_         field.Int64 // 设备组序号
	IP             field.String
	Status         field.Int64 // 0-正常；1-故障；2-报废

	fieldMap map[string]field.Expr
}

func (a avtEquipment) Table(newTableName string) *avtEquipment {
	a.avtEquipmentDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a avtEquipment) As(alias string) *avtEquipment {
	a.avtEquipmentDo.DO = *(a.avtEquipmentDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *avtEquipment) updateTableName(table string) *avtEquipment {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.Type = field.NewInt64(table, "type")
	a.Name = field.NewString(table, "name")
	a.Brand = field.NewString(table, "brand")
	a.PurchasingDate = field.NewString(table, "purchasing_date")
	a.Model = field.NewString(table, "model")
	a.ServiceLife = field.NewInt64(table, "service_life")
	a.WorkingTime = field.NewInt64(table, "working_time")
	a.IsDelete = field.NewInt64(table, "is_delete")
	a.CreateTime = field.NewTime(table, "create_time")
	a.UpdateTime = field.NewTime(table, "update_time")
	a.BuildingID = field.NewInt64(table, "building_id")
	a.FloorID = field.NewInt64(table, "floor_id")
	a.RoomID = field.NewInt64(table, "room_id")
	a.Group_ = field.NewInt64(table, "group")
	a.IP = field.NewString(table, "ip")
	a.Status = field.NewInt64(table, "status")

	a.fillFieldMap()

	return a
}

func (a *avtEquipment) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *avtEquipment) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 17)
	a.fieldMap["id"] = a.ID
	a.fieldMap["type"] = a.Type
	a.fieldMap["name"] = a.Name
	a.fieldMap["brand"] = a.Brand
	a.fieldMap["purchasing_date"] = a.PurchasingDate
	a.fieldMap["model"] = a.Model
	a.fieldMap["service_life"] = a.ServiceLife
	a.fieldMap["working_time"] = a.WorkingTime
	a.fieldMap["is_delete"] = a.IsDelete
	a.fieldMap["create_time"] = a.CreateTime
	a.fieldMap["update_time"] = a.UpdateTime
	a.fieldMap["building_id"] = a.BuildingID
	a.fieldMap["floor_id"] = a.FloorID
	a.fieldMap["room_id"] = a.RoomID
	a.fieldMap["group"] = a.Group_
	a.fieldMap["ip"] = a.IP
	a.fieldMap["status"] = a.Status
}

func (a avtEquipment) clone(db *gorm.DB) avtEquipment {
	a.avtEquipmentDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a avtEquipment) replaceDB(db *gorm.DB) avtEquipment {
	a.avtEquipmentDo.ReplaceDB(db)
	return a
}

type avtEquipmentDo struct{ gen.DO }

type IAvtEquipmentDo interface {
	gen.SubQuery
	Debug() IAvtEquipmentDo
	WithContext(ctx context.Context) IAvtEquipmentDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAvtEquipmentDo
	WriteDB() IAvtEquipmentDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAvtEquipmentDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAvtEquipmentDo
	Not(conds ...gen.Condition) IAvtEquipmentDo
	Or(conds ...gen.Condition) IAvtEquipmentDo
	Select(conds ...field.Expr) IAvtEquipmentDo
	Where(conds ...gen.Condition) IAvtEquipmentDo
	Order(conds ...field.Expr) IAvtEquipmentDo
	Distinct(cols ...field.Expr) IAvtEquipmentDo
	Omit(cols ...field.Expr) IAvtEquipmentDo
	Join(table schema.Tabler, on ...field.Expr) IAvtEquipmentDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAvtEquipmentDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAvtEquipmentDo
	Group(cols ...field.Expr) IAvtEquipmentDo
	Having(conds ...gen.Condition) IAvtEquipmentDo
	Limit(limit int) IAvtEquipmentDo
	Offset(offset int) IAvtEquipmentDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAvtEquipmentDo
	Unscoped() IAvtEquipmentDo
	Create(values ...*model.AvtEquipment) error
	CreateInBatches(values []*model.AvtEquipment, batchSize int) error
	Save(values ...*model.AvtEquipment) error
	First() (*model.AvtEquipment, error)
	Take() (*model.AvtEquipment, error)
	Last() (*model.AvtEquipment, error)
	Find() ([]*model.AvtEquipment, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AvtEquipment, err error)
	FindInBatches(result *[]*model.AvtEquipment, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.AvtEquipment) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAvtEquipmentDo
	Assign(attrs ...field.AssignExpr) IAvtEquipmentDo
	Joins(fields ...field.RelationField) IAvtEquipmentDo
	Preload(fields ...field.RelationField) IAvtEquipmentDo
	FirstOrInit() (*model.AvtEquipment, error)
	FirstOrCreate() (*model.AvtEquipment, error)
	FindByPage(offset int, limit int) (result []*model.AvtEquipment, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAvtEquipmentDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a avtEquipmentDo) Debug() IAvtEquipmentDo {
	return a.withDO(a.DO.Debug())
}

func (a avtEquipmentDo) WithContext(ctx context.Context) IAvtEquipmentDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a avtEquipmentDo) ReadDB() IAvtEquipmentDo {
	return a.Clauses(dbresolver.Read)
}

func (a avtEquipmentDo) WriteDB() IAvtEquipmentDo {
	return a.Clauses(dbresolver.Write)
}

func (a avtEquipmentDo) Session(config *gorm.Session) IAvtEquipmentDo {
	return a.withDO(a.DO.Session(config))
}

func (a avtEquipmentDo) Clauses(conds ...clause.Expression) IAvtEquipmentDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a avtEquipmentDo) Returning(value interface{}, columns ...string) IAvtEquipmentDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a avtEquipmentDo) Not(conds ...gen.Condition) IAvtEquipmentDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a avtEquipmentDo) Or(conds ...gen.Condition) IAvtEquipmentDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a avtEquipmentDo) Select(conds ...field.Expr) IAvtEquipmentDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a avtEquipmentDo) Where(conds ...gen.Condition) IAvtEquipmentDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a avtEquipmentDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IAvtEquipmentDo {
	return a.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (a avtEquipmentDo) Order(conds ...field.Expr) IAvtEquipmentDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a avtEquipmentDo) Distinct(cols ...field.Expr) IAvtEquipmentDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a avtEquipmentDo) Omit(cols ...field.Expr) IAvtEquipmentDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a avtEquipmentDo) Join(table schema.Tabler, on ...field.Expr) IAvtEquipmentDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a avtEquipmentDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAvtEquipmentDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a avtEquipmentDo) RightJoin(table schema.Tabler, on ...field.Expr) IAvtEquipmentDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a avtEquipmentDo) Group(cols ...field.Expr) IAvtEquipmentDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a avtEquipmentDo) Having(conds ...gen.Condition) IAvtEquipmentDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a avtEquipmentDo) Limit(limit int) IAvtEquipmentDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a avtEquipmentDo) Offset(offset int) IAvtEquipmentDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a avtEquipmentDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAvtEquipmentDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a avtEquipmentDo) Unscoped() IAvtEquipmentDo {
	return a.withDO(a.DO.Unscoped())
}

func (a avtEquipmentDo) Create(values ...*model.AvtEquipment) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a avtEquipmentDo) CreateInBatches(values []*model.AvtEquipment, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a avtEquipmentDo) Save(values ...*model.AvtEquipment) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a avtEquipmentDo) First() (*model.AvtEquipment, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AvtEquipment), nil
	}
}

func (a avtEquipmentDo) Take() (*model.AvtEquipment, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AvtEquipment), nil
	}
}

func (a avtEquipmentDo) Last() (*model.AvtEquipment, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AvtEquipment), nil
	}
}

func (a avtEquipmentDo) Find() ([]*model.AvtEquipment, error) {
	result, err := a.DO.Find()
	return result.([]*model.AvtEquipment), err
}

func (a avtEquipmentDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AvtEquipment, err error) {
	buf := make([]*model.AvtEquipment, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a avtEquipmentDo) FindInBatches(result *[]*model.AvtEquipment, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a avtEquipmentDo) Attrs(attrs ...field.AssignExpr) IAvtEquipmentDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a avtEquipmentDo) Assign(attrs ...field.AssignExpr) IAvtEquipmentDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a avtEquipmentDo) Joins(fields ...field.RelationField) IAvtEquipmentDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a avtEquipmentDo) Preload(fields ...field.RelationField) IAvtEquipmentDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a avtEquipmentDo) FirstOrInit() (*model.AvtEquipment, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AvtEquipment), nil
	}
}

func (a avtEquipmentDo) FirstOrCreate() (*model.AvtEquipment, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AvtEquipment), nil
	}
}

func (a avtEquipmentDo) FindByPage(offset int, limit int) (result []*model.AvtEquipment, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a avtEquipmentDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a avtEquipmentDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a avtEquipmentDo) Delete(models ...*model.AvtEquipment) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *avtEquipmentDo) withDO(do gen.Dao) *avtEquipmentDo {
	a.DO = *do.(*gen.DO)
	return a
}