package dto

import (
	"database/sql"
	"time"
)

type ProductResponse struct {
	ID          int64        `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Rating      float64      `json:"rating"`
	Image       string       `json:"image"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	DeletedAt   sql.NullTime `json:"deleted_at"`
}
