package repository

import (
	"context"
	"ecomtest/domain"

	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProduct(db *gorm.DB) domain.ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (p productRepository) FindByTitle(ctx context.Context, title string) (product domain.Product, err error) {
	err = p.db.Debug().WithContext(ctx).Table("prodcuts").First(&product, "title=?", title).Error
	return
}

func (p productRepository) Delete(ctx context.Context, id int64) error {
	err := p.db.Debug().WithContext(ctx).Table("products").Delete(&domain.Product{}, "id=?", id).Error
	return err
}

func (p productRepository) FindByID(ctx context.Context, id int64) (product domain.Product, err error) {
	err = p.db.Debug().WithContext(ctx).Table("products").First(&product, "id=?", id).Error
	return
}

func (p productRepository) GetAll(ctx context.Context) (products []domain.Product, err error) {
	err = p.db.Debug().WithContext(ctx).Table("products").Order("rating desc").Find(&products).Error
	return
}

func (p productRepository) Insert(ctx context.Context, product *domain.Product) error {
	err := p.db.Debug().WithContext(ctx).Table("products").Save(product).Error
	return err
}

func (p productRepository) Update(ctx context.Context, product *domain.Product) error {
	err := p.db.Debug().WithContext(ctx).Table("products").Updates(&product).Error
	return err
}
