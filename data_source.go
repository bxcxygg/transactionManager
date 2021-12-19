package transactionManager

import (
	"database/sql/driver"
)

// DataSource : 事务管理器的数据源
//
// 数据源用于向事务管理器注册事务数据源。
// 数据源可以开启事务
type DataSource interface {
	Name() string
	driver.ConnBeginTx
}
