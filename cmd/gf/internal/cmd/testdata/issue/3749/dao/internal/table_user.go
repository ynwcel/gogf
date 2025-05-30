// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TableUserDao is the data access object for the table table_user.
type TableUserDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  TableUserColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// TableUserColumns defines and stores column names for the table table_user.
type TableUserColumns struct {
	Id        string // User ID
	ParentId  string //
	Passport  string // User Passport
	PassWord  string // User Password
	Nickname2 string // User Nickname
	CreateAt  string // Created Time
	UpdateAt  string // Updated Time
}

// tableUserColumns holds the columns for the table table_user.
var tableUserColumns = TableUserColumns{
	Id:        "Id",
	ParentId:  "parentId",
	Passport:  "PASSPORT",
	PassWord:  "PASS_WORD",
	Nickname2: "NICKNAME2",
	CreateAt:  "create_at",
	UpdateAt:  "update_at",
}

// NewTableUserDao creates and returns a new DAO object for table data access.
func NewTableUserDao(handlers ...gdb.ModelHandler) *TableUserDao {
	return &TableUserDao{
		group:    "test",
		table:    "table_user",
		columns:  tableUserColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TableUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TableUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TableUserDao) Columns() TableUserColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TableUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TableUserDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *TableUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
