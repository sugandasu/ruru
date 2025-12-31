package nibirudb

import (
	"context"

	"gorm.io/gorm"
)

func (db *database) DB(ctx context.Context) *gorm.DB {
	txCtx := db.tx.GetContext(ctx)
	if txCtx == nil {
		return db.db.WithContext(ctx)
	}

	if txCtx.Progress == TransactionProgressNone {
		return db.db.WithContext(ctx)
	}

	if txCtx.Progress == TransactionProgressBegin {
		ses := db.db.Session(&gorm.Session{SkipDefaultTransaction: true})
		txCtx.Tx = ses.WithContext(ctx).Begin()
		txCtx.Progress = TransactionProgressRunning
	}

	txCtx.Count += 1

	return txCtx.Tx.WithContext(ctx)
}

func (db *database) Transaction() Transaction {
	return db.tx
}
