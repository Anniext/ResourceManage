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

func newAvtDeviceState(db *gorm.DB, opts ...gen.DOOption) avtDeviceState {
	_avtDeviceState := avtDeviceState{}

	_avtDeviceState.avtDeviceStateDo.UseDB(db, opts...)
	_avtDeviceState.avtDeviceStateDo.UseModel(&model.AvtDeviceState{})

	tableName := _avtDeviceState.avtDeviceStateDo.TableName()
	_avtDeviceState.ALL = field.NewAsterisk(tableName)
	_avtDeviceState.ID = field.NewInt64(tableName, "id")
	_avtDeviceState.RoomID = field.NewInt64(tableName, "room_id")
	_avtDeviceState.CreateTime = field.NewTime(tableName, "create_time")
	_avtDeviceState.UpdateTime = field.NewTime(tableName, "update_time")
	_avtDeviceState.Interval = field.NewInt64(tableName, "interval")
	_avtDeviceState.JoinNum = field.NewInt64(tableName, "join_num")

	_avtDeviceState.fillFieldMap()

	return _avtDeviceState
}

type avtDeviceState struct {
	avtDeviceStateDo

	ALL        field.Asterisk
	ID         field.Int64
	RoomID     field.Int64
	CreateTime field.Time  // 上课时间
	UpdateTime field.Time  // 下课时间
	Interval   field.Int64 // 时间间隔: 更新时间-创建时间
	JoinNum    field.Int64

	fieldMap map[string]field.Expr
}

func (a avtDeviceState) Table(newTableName string) *avtDeviceState {
	a.avtDeviceStateDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a avtDeviceState) As(alias string) *avtDeviceState {
	a.avtDeviceStateDo.DO = *(a.avtDeviceStateDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *avtDeviceState) updateTableName(table string) *avtDeviceState {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.RoomID = field.NewInt64(table, "room_id")
	a.CreateTime = field.NewTime(table, "create_time")
	a.UpdateTime = field.NewTime(table, "update_time")
	a.Interval = field.NewInt64(table, "interval")
	a.JoinNum = field.NewInt64(table, "join_num")

	a.fillFieldMap()

	return a
}

func (a *avtDeviceState) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *avtDeviceState) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 6)
	a.fieldMap["id"] = a.ID
	a.fieldMap["room_id"] = a.RoomID
	a.fieldMap["create_time"] = a.CreateTime
	a.fieldMap["update_time"] = a.UpdateTime
	a.fieldMap["interval"] = a.Interval
	a.fieldMap["join_num"] = a.JoinNum
}

func (a avtDeviceState) clone(db *gorm.DB) avtDeviceState {
	a.avtDeviceStateDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a avtDeviceState) replaceDB(db *gorm.DB) avtDeviceState {
	a.avtDeviceStateDo.ReplaceDB(db)
	return a
}

type avtDeviceStateDo struct{ gen.DO }

type IAvtDeviceStateDo interface {
	gen.SubQuery
	Debug() IAvtDeviceStateDo
	WithContext(ctx context.Context) IAvtDeviceStateDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAvtDeviceStateDo
	WriteDB() IAvtDeviceStateDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAvtDeviceStateDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAvtDeviceStateDo
	Not(conds ...gen.Condition) IAvtDeviceStateDo
	Or(conds ...gen.Condition) IAvtDeviceStateDo
	Select(conds ...field.Expr) IAvtDeviceStateDo
	Where(conds ...gen.Condition) IAvtDeviceStateDo
	Order(conds ...field.Expr) IAvtDeviceStateDo
	Distinct(cols ...field.Expr) IAvtDeviceStateDo
	Omit(cols ...field.Expr) IAvtDeviceStateDo
	Join(table schema.Tabler, on ...field.Expr) IAvtDeviceStateDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAvtDeviceStateDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAvtDeviceStateDo
	Group(cols ...field.Expr) IAvtDeviceStateDo
	Having(conds ...gen.Condition) IAvtDeviceStateDo
	Limit(limit int) IAvtDeviceStateDo
	Offset(offset int) IAvtDeviceStateDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAvtDeviceStateDo
	Unscoped() IAvtDeviceStateDo
	Create(values ...*model.AvtDeviceState) error
	CreateInBatches(values []*model.AvtDeviceState, batchSize int) error
	Save(values ...*model.AvtDeviceState) error
	First() (*model.AvtDeviceState, error)
	Take() (*model.AvtDeviceState, error)
	Last() (*model.AvtDeviceState, error)
	Find() ([]*model.AvtDeviceState, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AvtDeviceState, err error)
	FindInBatches(result *[]*model.AvtDeviceState, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.AvtDeviceState) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAvtDeviceStateDo
	Assign(attrs ...field.AssignExpr) IAvtDeviceStateDo
	Joins(fields ...field.RelationField) IAvtDeviceStateDo
	Preload(fields ...field.RelationField) IAvtDeviceStateDo
	FirstOrInit() (*model.AvtDeviceState, error)
	FirstOrCreate() (*model.AvtDeviceState, error)
	FindByPage(offset int, limit int) (result []*model.AvtDeviceState, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAvtDeviceStateDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a avtDeviceStateDo) Debug() IAvtDeviceStateDo {
	return a.withDO(a.DO.Debug())
}

func (a avtDeviceStateDo) WithContext(ctx context.Context) IAvtDeviceStateDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a avtDeviceStateDo) ReadDB() IAvtDeviceStateDo {
	return a.Clauses(dbresolver.Read)
}

func (a avtDeviceStateDo) WriteDB() IAvtDeviceStateDo {
	return a.Clauses(dbresolver.Write)
}

func (a avtDeviceStateDo) Session(config *gorm.Session) IAvtDeviceStateDo {
	return a.withDO(a.DO.Session(config))
}

func (a avtDeviceStateDo) Clauses(conds ...clause.Expression) IAvtDeviceStateDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a avtDeviceStateDo) Returning(value interface{}, columns ...string) IAvtDeviceStateDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a avtDeviceStateDo) Not(conds ...gen.Condition) IAvtDeviceStateDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a avtDeviceStateDo) Or(conds ...gen.Condition) IAvtDeviceStateDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a avtDeviceStateDo) Select(conds ...field.Expr) IAvtDeviceStateDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a avtDeviceStateDo) Where(conds ...gen.Condition) IAvtDeviceStateDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a avtDeviceStateDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IAvtDeviceStateDo {
	return a.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (a avtDeviceStateDo) Order(conds ...field.Expr) IAvtDeviceStateDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a avtDeviceStateDo) Distinct(cols ...field.Expr) IAvtDeviceStateDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a avtDeviceStateDo) Omit(cols ...field.Expr) IAvtDeviceStateDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a avtDeviceStateDo) Join(table schema.Tabler, on ...field.Expr) IAvtDeviceStateDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a avtDeviceStateDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAvtDeviceStateDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a avtDeviceStateDo) RightJoin(table schema.Tabler, on ...field.Expr) IAvtDeviceStateDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a avtDeviceStateDo) Group(cols ...field.Expr) IAvtDeviceStateDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a avtDeviceStateDo) Having(conds ...gen.Condition) IAvtDeviceStateDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a avtDeviceStateDo) Limit(limit int) IAvtDeviceStateDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a avtDeviceStateDo) Offset(offset int) IAvtDeviceStateDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a avtDeviceStateDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAvtDeviceStateDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a avtDeviceStateDo) Unscoped() IAvtDeviceStateDo {
	return a.withDO(a.DO.Unscoped())
}

func (a avtDeviceStateDo) Create(values ...*model.AvtDeviceState) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a avtDeviceStateDo) CreateInBatches(values []*model.AvtDeviceState, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a avtDeviceStateDo) Save(values ...*model.AvtDeviceState) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a avtDeviceStateDo) First() (*model.AvtDeviceState, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AvtDeviceState), nil
	}
}

func (a avtDeviceStateDo) Take() (*model.AvtDeviceState, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AvtDeviceState), nil
	}
}

func (a avtDeviceStateDo) Last() (*model.AvtDeviceState, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AvtDeviceState), nil
	}
}

func (a avtDeviceStateDo) Find() ([]*model.AvtDeviceState, error) {
	result, err := a.DO.Find()
	return result.([]*model.AvtDeviceState), err
}

func (a avtDeviceStateDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AvtDeviceState, err error) {
	buf := make([]*model.AvtDeviceState, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a avtDeviceStateDo) FindInBatches(result *[]*model.AvtDeviceState, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a avtDeviceStateDo) Attrs(attrs ...field.AssignExpr) IAvtDeviceStateDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a avtDeviceStateDo) Assign(attrs ...field.AssignExpr) IAvtDeviceStateDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a avtDeviceStateDo) Joins(fields ...field.RelationField) IAvtDeviceStateDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a avtDeviceStateDo) Preload(fields ...field.RelationField) IAvtDeviceStateDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a avtDeviceStateDo) FirstOrInit() (*model.AvtDeviceState, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AvtDeviceState), nil
	}
}

func (a avtDeviceStateDo) FirstOrCreate() (*model.AvtDeviceState, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AvtDeviceState), nil
	}
}

func (a avtDeviceStateDo) FindByPage(offset int, limit int) (result []*model.AvtDeviceState, count int64, err error) {
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

func (a avtDeviceStateDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a avtDeviceStateDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a avtDeviceStateDo) Delete(models ...*model.AvtDeviceState) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *avtDeviceStateDo) withDO(do gen.Dao) *avtDeviceStateDo {
	a.DO = *do.(*gen.DO)
	return a
}
