package domain

import (
	"context"
	"database/sql"
	"ecomtest/dto"
	"time"
)

type Product struct {
	ID          int64 `gorm:"not null; primaryKey; autoIncrement"`
	Title       string
	Description string
	Rating      float64
	Image       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   sql.NullTime
}

type ProductRepository interface {
	GetAll(ctx context.Context) ([]Product, error)
	FindByID(ctx context.Context, id int64) (Product, error)
	FindByTitle(ctx context.Context, title string) (Product, error)
	Insert(ctx context.Context, product *Product) error
	Update(ctx context.Context, product *Product) error
	Delete(ctx context.Context, id int64) error
}

type ProductService interface {
	ShowAllProducts(ctx context.Context) dto.Response
	GetDetailProduct(ctx context.Context, req int64) dto.Response
	CreateNewProduct(ctx context.Context, req dto.ProductRequest) dto.Response
	UpdateTheProduct(ctx context.Context, req dto.ProductRequest) dto.Response
	DeleteTheProduct(ctx context.Context, req int64) dto.Response
}
