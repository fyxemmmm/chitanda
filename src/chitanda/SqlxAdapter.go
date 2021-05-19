package chitanda

import (
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

func NewSqlXAdapter() *SqlXAdapter {
	db := sqlx.MustConnect("mysql",
		"root:pwd@tcp(121.37.156.155:3306)/")
	db.SetConnMaxLifetime(DefaultConMaxLifeTime)
	db.SetMaxIdleConns(DefaultMaxIdleConns)
	db.SetMaxOpenConns(DefaultMaxOpenConns)
	return &SqlXAdapter{db}
}
