package dbrepo

import "database/sql"

// Store 组合Querier接口进行扩展
type Store interface {
	Querier

	// 定义自己的方法
	// ...
}

// SQLStore 提供执行SQL查询和事务的所有功能, 实现 Store 接口
type SQLStore struct {
	db *sql.DB
	*Queries
}

// NewSQLStore 创建一个 SQLStore 实现 Store 接口，操作DB
func NewSQLStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}
