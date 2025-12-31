package nibirudb

import (
	"context"
	"log"

	"gorm.io/gorm"
)

type Transaction interface {
	ExecTransaction(ctx context.Context, fn TransactionFunction) (context.Context, error)
	GetContext(ctx context.Context) *TransactionContext
}

type transaction struct {
	Context *TransactionContext
}

type TransactionFunction func(ctx context.Context) (context.Context, error)

type TransactionProgress string

const (
	TransactionProgressNone    TransactionProgress = "NONE"
	TransactionProgressBegin   TransactionProgress = "BEGIN"
	TransactionProgressRunning TransactionProgress = "RUNNING"
)

type TransactionContext struct {
	Progress TransactionProgress
	Count    int
	Tx       *gorm.DB
}

type TransactionKey struct{}

func NewTransaction(db *gorm.DB) Transaction {
	ctx := &TransactionContext{
		Progress: TransactionProgressNone,
		Count:    0,
		Tx:       db,
	}
	return &transaction{
		Context: ctx,
	}
}

func (tx *transaction) ExecTransaction(ctx context.Context, fn TransactionFunction) (context.Context, error) {
	txCtx := &TransactionContext{
		Progress: TransactionProgressBegin,
		Count:    0,
		Tx:       &gorm.DB{},
	}
	ctx = context.WithValue(ctx, TransactionKey{}, txCtx)

	ctx, err := fn(ctx)

	if err != nil {
		txErr := txCtx.Tx.Rollback().Error
		if txErr != nil {
			log.Println("Database:", txErr)
		}
		ctx = context.WithValue(ctx, TransactionKey{}, nil)
		return ctx, err
	}

	txErr := txCtx.Tx.Commit().Error
	if txErr != nil {
		log.Println("Database:", txErr)
	}

	ctx = context.WithValue(ctx, TransactionKey{}, nil)
	return ctx, err
}

func (tx *transaction) GetContext(ctx context.Context) *TransactionContext {
	val := ctx.Value(TransactionKey{})
	if val == nil {
		return nil
	}

	txCtx, ok := val.(*TransactionContext)
	if !ok {
		return nil
	}

	return txCtx
}
