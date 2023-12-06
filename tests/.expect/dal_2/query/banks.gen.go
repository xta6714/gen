// Code generated by github.com/xta6714/gen. DO NOT EDIT.
// Code generated by github.com/xta6714/gen. DO NOT EDIT.
// Code generated by github.com/xta6714/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"github.com/xta6714/gen"
	"github.com/xta6714/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/xta6714/gen/tests/.gen/dal_2/model"
)

func newBank(db *gorm.DB, opts ...gen.DOOption) bank {
	_bank := bank{}

	_bank.bankDo.UseDB(db, opts...)
	_bank.bankDo.UseModel(&model.Bank{})

	tableName := _bank.bankDo.TableName()
	_bank.ALL = field.NewAsterisk(tableName)
	_bank.ID = field.NewInt64(tableName, "id")
	_bank.Name = field.NewString(tableName, "name")
	_bank.Address = field.NewString(tableName, "address")
	_bank.Scale = field.NewInt64(tableName, "scale")

	_bank.fillFieldMap()

	return _bank
}

type bank struct {
	bankDo bankDo

	ALL     field.Asterisk
	ID      field.Int64
	Name    field.String
	Address field.String
	Scale   field.Int64

	fieldMap map[string]field.Expr
}

func (b bank) Table(newTableName string) *bank {
	b.bankDo.UseTable(newTableName)
	return b.updateTableName(newTableName)
}

func (b bank) As(alias string) *bank {
	b.bankDo.DO = *(b.bankDo.As(alias).(*gen.DO))
	return b.updateTableName(alias)
}

func (b *bank) updateTableName(table string) *bank {
	b.ALL = field.NewAsterisk(table)
	b.ID = field.NewInt64(table, "id")
	b.Name = field.NewString(table, "name")
	b.Address = field.NewString(table, "address")
	b.Scale = field.NewInt64(table, "scale")

	b.fillFieldMap()

	return b
}

func (b *bank) WithContext(ctx context.Context) *bankDo { return b.bankDo.WithContext(ctx) }

func (b bank) TableName() string { return b.bankDo.TableName() }

func (b bank) Alias() string { return b.bankDo.Alias() }

func (b bank) Columns(cols ...field.Expr) gen.Columns { return b.bankDo.Columns(cols...) }

func (b *bank) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := b.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (b *bank) fillFieldMap() {
	b.fieldMap = make(map[string]field.Expr, 4)
	b.fieldMap["id"] = b.ID
	b.fieldMap["name"] = b.Name
	b.fieldMap["address"] = b.Address
	b.fieldMap["scale"] = b.Scale
}

func (b bank) clone(db *gorm.DB) bank {
	b.bankDo.ReplaceConnPool(db.Statement.ConnPool)
	return b
}

func (b bank) replaceDB(db *gorm.DB) bank {
	b.bankDo.ReplaceDB(db)
	return b
}

type bankDo struct{ gen.DO }

func (b bankDo) Debug() *bankDo {
	return b.withDO(b.DO.Debug())
}

func (b bankDo) WithContext(ctx context.Context) *bankDo {
	return b.withDO(b.DO.WithContext(ctx))
}

func (b bankDo) ReadDB() *bankDo {
	return b.Clauses(dbresolver.Read)
}

func (b bankDo) WriteDB() *bankDo {
	return b.Clauses(dbresolver.Write)
}

func (b bankDo) Session(config *gorm.Session) *bankDo {
	return b.withDO(b.DO.Session(config))
}

func (b bankDo) Clauses(conds ...clause.Expression) *bankDo {
	return b.withDO(b.DO.Clauses(conds...))
}

func (b bankDo) Returning(value interface{}, columns ...string) *bankDo {
	return b.withDO(b.DO.Returning(value, columns...))
}

func (b bankDo) Not(conds ...gen.Condition) *bankDo {
	return b.withDO(b.DO.Not(conds...))
}

func (b bankDo) Or(conds ...gen.Condition) *bankDo {
	return b.withDO(b.DO.Or(conds...))
}

func (b bankDo) Select(conds ...field.Expr) *bankDo {
	return b.withDO(b.DO.Select(conds...))
}

func (b bankDo) Where(conds ...gen.Condition) *bankDo {
	return b.withDO(b.DO.Where(conds...))
}

func (b bankDo) Order(conds ...field.Expr) *bankDo {
	return b.withDO(b.DO.Order(conds...))
}

func (b bankDo) Distinct(cols ...field.Expr) *bankDo {
	return b.withDO(b.DO.Distinct(cols...))
}

func (b bankDo) Omit(cols ...field.Expr) *bankDo {
	return b.withDO(b.DO.Omit(cols...))
}

func (b bankDo) Join(table schema.Tabler, on ...field.Expr) *bankDo {
	return b.withDO(b.DO.Join(table, on...))
}

func (b bankDo) LeftJoin(table schema.Tabler, on ...field.Expr) *bankDo {
	return b.withDO(b.DO.LeftJoin(table, on...))
}

func (b bankDo) RightJoin(table schema.Tabler, on ...field.Expr) *bankDo {
	return b.withDO(b.DO.RightJoin(table, on...))
}

func (b bankDo) Group(cols ...field.Expr) *bankDo {
	return b.withDO(b.DO.Group(cols...))
}

func (b bankDo) Having(conds ...gen.Condition) *bankDo {
	return b.withDO(b.DO.Having(conds...))
}

func (b bankDo) Limit(limit int) *bankDo {
	return b.withDO(b.DO.Limit(limit))
}

func (b bankDo) Offset(offset int) *bankDo {
	return b.withDO(b.DO.Offset(offset))
}

func (b bankDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *bankDo {
	return b.withDO(b.DO.Scopes(funcs...))
}

func (b bankDo) Unscoped() *bankDo {
	return b.withDO(b.DO.Unscoped())
}

func (b bankDo) Create(values ...*model.Bank) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Create(values)
}

func (b bankDo) CreateInBatches(values []*model.Bank, batchSize int) error {
	return b.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (b bankDo) Save(values ...*model.Bank) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Save(values)
}

func (b bankDo) First() (*model.Bank, error) {
	if result, err := b.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Bank), nil
	}
}

func (b bankDo) Take() (*model.Bank, error) {
	if result, err := b.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Bank), nil
	}
}

func (b bankDo) Last() (*model.Bank, error) {
	if result, err := b.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Bank), nil
	}
}

func (b bankDo) Find() ([]*model.Bank, error) {
	result, err := b.DO.Find()
	return result.([]*model.Bank), err
}

func (b bankDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Bank, err error) {
	buf := make([]*model.Bank, 0, batchSize)
	err = b.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (b bankDo) FindInBatches(result *[]*model.Bank, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return b.DO.FindInBatches(result, batchSize, fc)
}

func (b bankDo) Attrs(attrs ...field.AssignExpr) *bankDo {
	return b.withDO(b.DO.Attrs(attrs...))
}

func (b bankDo) Assign(attrs ...field.AssignExpr) *bankDo {
	return b.withDO(b.DO.Assign(attrs...))
}

func (b bankDo) Joins(fields ...field.RelationField) *bankDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Joins(_f))
	}
	return &b
}

func (b bankDo) Preload(fields ...field.RelationField) *bankDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Preload(_f))
	}
	return &b
}

func (b bankDo) FirstOrInit() (*model.Bank, error) {
	if result, err := b.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Bank), nil
	}
}

func (b bankDo) FirstOrCreate() (*model.Bank, error) {
	if result, err := b.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Bank), nil
	}
}

func (b bankDo) FindByPage(offset int, limit int) (result []*model.Bank, count int64, err error) {
	result, err = b.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = b.Offset(-1).Limit(-1).Count()
	return
}

func (b bankDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = b.Count()
	if err != nil {
		return
	}

	err = b.Offset(offset).Limit(limit).Scan(result)
	return
}

func (b bankDo) Scan(result interface{}) (err error) {
	return b.DO.Scan(result)
}

func (b bankDo) Delete(models ...*model.Bank) (result gen.ResultInfo, err error) {
	return b.DO.Delete(models)
}

func (b *bankDo) withDO(do gen.Dao) *bankDo {
	b.DO = *do.(*gen.DO)
	return b
}
