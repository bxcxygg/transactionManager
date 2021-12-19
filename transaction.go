package transactionManager

import (
	"context"
	"database/sql/driver"
)

type UseCase interface {
	UseTx(tx driver.Tx) error
}

type Transaction interface {
	driver.Tx
	Build(...UseCase) Transaction
	TxEnd(func() error) error
}

// TransactionManager : 事务管理器
//
// 事务管理器用于管理事务数据源
type TransactionManager interface {
	Register(source DataSource)
	EnableTx(ctx context.Context, opts driver.TxOptions) (Transaction, error)
}

type transactionManager struct {
	dataSource DataSource
}

func (t *transactionManager) Register(source DataSource) {
	t.dataSource = source
}

func (t *transactionManager) EnableTx(ctx context.Context, opts driver.TxOptions) (Transaction, error) {
	return nil, nil
}

type transaction struct {
	err error
	tx  driver.Tx
}

func (t *transaction) Commit() error {
	return t.tx.Commit()
}

func (t *transaction) Rollback() error {
	return t.tx.Rollback()
}

func (t *transaction) Build(useCase ...UseCase) Transaction {
	if t.err != nil {
		return t
	}
	for _, uc := range useCase {
		err := uc.UseTx(t.tx)
		if err != nil {
			t.err = err
			return t
		}
	}
	return t
}

func (t *transaction) TxEnd(f func() error) error {
	if t.err != nil {
		return t.err
	}
	return f()
}
