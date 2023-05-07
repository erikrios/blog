package config

import (
	"database/sql"
	"fmt"

	"github.com/erikrios/blog/constant"
	_ "github.com/lib/pq"
)

func NewPostgreSQLDatabase() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		constant.DBHost,
		constant.DBPort,
		constant.DBUser,
		constant.DBPassword,
		constant.DBName,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	return db, nil
}
