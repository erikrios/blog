package entity

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int64
	Username  string
	Name      string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}
