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

func newSysOperator(db *gorm.DB, opts ...gen.DOOption) sysOperator {
	_sysOperator := sysOperator{}

	_sysOperator.sysOperatorDo.UseDB(db, opts...)
	_sysOperator.sysOperatorDo.UseModel(&model.SysOperator{})

	tableName := _sysOperator.sysOperatorDo.TableName()
	_sysOperator.ALL = field.NewAsterisk(tableName)
	_sysOperator.ID = field.NewInt64(tableName, "id")
	_sysOperator.OperatorNo = field.NewInt64(tableName, "operator_No")
	_sysOperator.Name = field.NewString(tableName, "name")
	_sysOperator.OperatorDes = field.NewString(tableName, "operator_des")
	_sysOperator.Contact = field.NewString(tableName, "contact")
	_sysOperator.Address = field.NewString(tableName, "address")
	_sysOperator.Status = field.NewInt64(tableName, "status")
	_sysOperator.CreateTime = field.NewTime(tableName, "create_time")
	_sysOperator.UpdateTime = field.NewTime(tableName, "update_time")
	_sysOperator.OperatorType = field.NewInt64(tableName, "operator_type")
	_sysOperator.Mobile = field.NewString(tableName, "mobile")
	_sysOperator.Email = field.NewString(tableName, "email")
	_sysOperator.Discount = field.NewFloat32(tableName, "discount")
	_sysOperator.Earnings = field.NewInt64(tableName, "earnings")
	_sysOperator.DeleteFlag = field.NewInt64(tableName, "delete_flag")

	_sysOperator.fillFieldMap()

	return _sysOperator
}

type sysOperator struct {
	sysOperatorDo

	ALL          field.Asterisk
	ID           field.Int64
	OperatorNo   field.Int64
	Name         field.String // 运营商名称
	OperatorDes  field.String // 运营商详情
	Contact      field.String // 联系人
	Address      field.String
	Status       field.Int64 // 0,无效的；1，有效的
	CreateTime   field.Time
	UpdateTime   field.Time
	OperatorType field.Int64 // 2-运营商；1-大客户;3-推广渠道；4-协议企业
	Mobile       field.String
	Email        field.String
	Discount     field.Float32 // 对于推广渠道，表示服务费分成比例；对于协议公司表示服务费优惠比例
	Earnings     field.Int64   // 收入分成
	DeleteFlag   field.Int64   // 0-为删除；1-已删除

	fieldMap map[string]field.Expr
}

func (s sysOperator) Table(newTableName string) *sysOperator {
	s.sysOperatorDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysOperator) As(alias string) *sysOperator {
	s.sysOperatorDo.DO = *(s.sysOperatorDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysOperator) updateTableName(table string) *sysOperator {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.OperatorNo = field.NewInt64(table, "operator_No")
	s.Name = field.NewString(table, "name")
	s.OperatorDes = field.NewString(table, "operator_des")
	s.Contact = field.NewString(table, "contact")
	s.Address = field.NewString(table, "address")
	s.Status = field.NewInt64(table, "status")
	s.CreateTime = field.NewTime(table, "create_time")
	s.UpdateTime = field.NewTime(table, "update_time")
	s.OperatorType = field.NewInt64(table, "operator_type")
	s.Mobile = field.NewString(table, "mobile")
	s.Email = field.NewString(table, "email")
	s.Discount = field.NewFloat32(table, "discount")
	s.Earnings = field.NewInt64(table, "earnings")
	s.DeleteFlag = field.NewInt64(table, "delete_flag")

	s.fillFieldMap()

	return s
}

func (s *sysOperator) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysOperator) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 15)
	s.fieldMap["id"] = s.ID
	s.fieldMap["operator_No"] = s.OperatorNo
	s.fieldMap["name"] = s.Name
	s.fieldMap["operator_des"] = s.OperatorDes
	s.fieldMap["contact"] = s.Contact
	s.fieldMap["address"] = s.Address
	s.fieldMap["status"] = s.Status
	s.fieldMap["create_time"] = s.CreateTime
	s.fieldMap["update_time"] = s.UpdateTime
	s.fieldMap["operator_type"] = s.OperatorType
	s.fieldMap["mobile"] = s.Mobile
	s.fieldMap["email"] = s.Email
	s.fieldMap["discount"] = s.Discount
	s.fieldMap["earnings"] = s.Earnings
	s.fieldMap["delete_flag"] = s.DeleteFlag
}

func (s sysOperator) clone(db *gorm.DB) sysOperator {
	s.sysOperatorDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysOperator) replaceDB(db *gorm.DB) sysOperator {
	s.sysOperatorDo.ReplaceDB(db)
	return s
}

type sysOperatorDo struct{ gen.DO }

type ISysOperatorDo interface {
	gen.SubQuery
	Debug() ISysOperatorDo
	WithContext(ctx context.Context) ISysOperatorDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISysOperatorDo
	WriteDB() ISysOperatorDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISysOperatorDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISysOperatorDo
	Not(conds ...gen.Condition) ISysOperatorDo
	Or(conds ...gen.Condition) ISysOperatorDo
	Select(conds ...field.Expr) ISysOperatorDo
	Where(conds ...gen.Condition) ISysOperatorDo
	Order(conds ...field.Expr) ISysOperatorDo
	Distinct(cols ...field.Expr) ISysOperatorDo
	Omit(cols ...field.Expr) ISysOperatorDo
	Join(table schema.Tabler, on ...field.Expr) ISysOperatorDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISysOperatorDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISysOperatorDo
	Group(cols ...field.Expr) ISysOperatorDo
	Having(conds ...gen.Condition) ISysOperatorDo
	Limit(limit int) ISysOperatorDo
	Offset(offset int) ISysOperatorDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISysOperatorDo
	Unscoped() ISysOperatorDo
	Create(values ...*model.SysOperator) error
	CreateInBatches(values []*model.SysOperator, batchSize int) error
	Save(values ...*model.SysOperator) error
	First() (*model.SysOperator, error)
	Take() (*model.SysOperator, error)
	Last() (*model.SysOperator, error)
	Find() ([]*model.SysOperator, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysOperator, err error)
	FindInBatches(result *[]*model.SysOperator, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SysOperator) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISysOperatorDo
	Assign(attrs ...field.AssignExpr) ISysOperatorDo
	Joins(fields ...field.RelationField) ISysOperatorDo
	Preload(fields ...field.RelationField) ISysOperatorDo
	FirstOrInit() (*model.SysOperator, error)
	FirstOrCreate() (*model.SysOperator, error)
	FindByPage(offset int, limit int) (result []*model.SysOperator, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISysOperatorDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sysOperatorDo) Debug() ISysOperatorDo {
	return s.withDO(s.DO.Debug())
}

func (s sysOperatorDo) WithContext(ctx context.Context) ISysOperatorDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysOperatorDo) ReadDB() ISysOperatorDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysOperatorDo) WriteDB() ISysOperatorDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysOperatorDo) Session(config *gorm.Session) ISysOperatorDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysOperatorDo) Clauses(conds ...clause.Expression) ISysOperatorDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysOperatorDo) Returning(value interface{}, columns ...string) ISysOperatorDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysOperatorDo) Not(conds ...gen.Condition) ISysOperatorDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysOperatorDo) Or(conds ...gen.Condition) ISysOperatorDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysOperatorDo) Select(conds ...field.Expr) ISysOperatorDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysOperatorDo) Where(conds ...gen.Condition) ISysOperatorDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysOperatorDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ISysOperatorDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s sysOperatorDo) Order(conds ...field.Expr) ISysOperatorDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysOperatorDo) Distinct(cols ...field.Expr) ISysOperatorDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysOperatorDo) Omit(cols ...field.Expr) ISysOperatorDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysOperatorDo) Join(table schema.Tabler, on ...field.Expr) ISysOperatorDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysOperatorDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISysOperatorDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysOperatorDo) RightJoin(table schema.Tabler, on ...field.Expr) ISysOperatorDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysOperatorDo) Group(cols ...field.Expr) ISysOperatorDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysOperatorDo) Having(conds ...gen.Condition) ISysOperatorDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysOperatorDo) Limit(limit int) ISysOperatorDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysOperatorDo) Offset(offset int) ISysOperatorDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysOperatorDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISysOperatorDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysOperatorDo) Unscoped() ISysOperatorDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysOperatorDo) Create(values ...*model.SysOperator) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysOperatorDo) CreateInBatches(values []*model.SysOperator, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysOperatorDo) Save(values ...*model.SysOperator) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysOperatorDo) First() (*model.SysOperator, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysOperator), nil
	}
}

func (s sysOperatorDo) Take() (*model.SysOperator, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysOperator), nil
	}
}

func (s sysOperatorDo) Last() (*model.SysOperator, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysOperator), nil
	}
}

func (s sysOperatorDo) Find() ([]*model.SysOperator, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysOperator), err
}

func (s sysOperatorDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysOperator, err error) {
	buf := make([]*model.SysOperator, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysOperatorDo) FindInBatches(result *[]*model.SysOperator, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysOperatorDo) Attrs(attrs ...field.AssignExpr) ISysOperatorDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysOperatorDo) Assign(attrs ...field.AssignExpr) ISysOperatorDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysOperatorDo) Joins(fields ...field.RelationField) ISysOperatorDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysOperatorDo) Preload(fields ...field.RelationField) ISysOperatorDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysOperatorDo) FirstOrInit() (*model.SysOperator, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysOperator), nil
	}
}

func (s sysOperatorDo) FirstOrCreate() (*model.SysOperator, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysOperator), nil
	}
}

func (s sysOperatorDo) FindByPage(offset int, limit int) (result []*model.SysOperator, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s sysOperatorDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysOperatorDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysOperatorDo) Delete(models ...*model.SysOperator) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysOperatorDo) withDO(do gen.Dao) *sysOperatorDo {
	s.DO = *do.(*gen.DO)
	return s
}