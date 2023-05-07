package entity

import (
	"database/sql"
	"time"
)

type Article struct {
	ID         int64
	UserID     int64
	CategoryID int64
	Title      string
	Content    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  sql.NullTime
}
