package chitanda

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

const (
	DefaultConMaxLifeTime = 30 * time.Minute
	DefaultMaxIdleConns   = 10
	DefaultMaxOpenConns   = 50
)

type SqlXAdapter struct {
	*sqlx.DB
}

func NewSqlXAdapter(host string, username string, password string) *SqlXAdapter {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/", username, password, host)
	db := sqlx.MustConnect("mysql", dsn)
	db.SetConnMaxLifetime(DefaultConMaxLifeTime)
	db.SetMaxIdleConns(DefaultMaxIdleConns)
	db.SetMaxOpenConns(DefaultMaxOpenConns)
	return &SqlXAdapter{db}
}
