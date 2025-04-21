package db

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var pool *sql.DB

// Init creates a global connection pool to viewdb
func Init(user, pass, host string, port int) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/db?parseTime=true&loc=Asia%%2FTokyo",
		user, pass, host, port)

	var err error
	pool, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// Optionally configure pool sizes here…
	return pool.Ping()
}

// ConnForTenant sets @current_tenant_id on a fresh connection.
func ConnForTenant(ctx context.Context, tenantID uint64) (*sql.DB, error) {
	if _, err := pool.ExecContext(ctx, "SET @current_tenant_id = ?", tenantID); err != nil {
		return nil, err
	}
	return pool, nil
}
