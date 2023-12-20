package service

import (
	"context"
	"ecomtest/domain"
	"ecomtest/dto"
	"time"
)

type productService struct {
	productRepository domain.ProductRepository
}

func NewProduct(productRepository domain.ProductRepository) domain.ProductService {
	return &productService{
		productRepository: productRepository,
	}
}

func (p productService) CreateNewProduct(ctx context.Context, req dto.ProductRequest) dto.Response {
	if req == (dto.ProductRequest{}) {
		return dto.Response{
			Code:    400,
			Message: "INVALID",
			Error:   "field must be fill",
		}
	}
	exist, _ := p.productRepository.FindByTitle(ctx, req.Title)
	if exist != (domain.Product{}) {
		return dto.Response{
			Code:    401,
			Message: "INVALID",
			Error:   "title must be unique",
		}
	}
	var product = domain.Product{
		Title:       req.Title,
		Description: req.Description,
		Rating:      req.Rating,
		Image:       req.Image,
		CreatedAt:   time.Now(),
	}
	if err := p.productRepository.Insert(ctx, &product); err != nil {
		return dto.Response{
			Code:    401,
			Message: "INVALID",
			Error:   err.Error(),
		}
	}

	return dto.Response{
		Code:    200,
		Message: "APPROVE",
	}
}

func (p productService) DeleteTheProduct(ctx context.Context, req int64) dto.Response {
	if err := p.productRepository.Delete(ctx, req); err != nil {
		return dto.Response{
			Code:    400,
			Message: "INVALID",
			Error:   err.Error(),
		}
	}
	return dto.Response{
		Code:    200,
		Message: "APPROVE",
	}
}

func (p productService) GetDetailProduct(ctx context.Context, req int64) dto.Response {
	product, err := p.productRepository.FindByID(ctx, req)
	if err != nil {
		return dto.Response{
			Code:    404,
			Message: "INVALID",
			Error:   err.Error(),
		}
	}

	return dto.Response{
		Code:    200,
		Message: "APPROVE",
		Data: dto.ProductResponse{
			ID:          product.ID,
			Title:       product.Title,
			Description: product.Description,
			Rating:      product.Rating,
			Image:       product.Image,
			CreatedAt:   product.CreatedAt,
			UpdatedAt:   product.UpdatedAt,
			DeletedAt:   product.DeletedAt,
		},
	}

}

func (p productService) ShowAllProducts(ctx context.Context) dto.Response {
	products, err := p.productRepository.GetAll(ctx)
	if err != nil {
		return dto.Response{
			Code:    400,
			Message: "INVALID",
			Error:   err.Error(),
		}
	}

	var productsResponse []dto.ProductResponse

	for _, v := range products {
		productsResponse = append(productsResponse, dto.ProductResponse{
			ID:          v.ID,
			Title:       v.Title,
			Description: v.Description,
			Rating:      v.Rating,
			Image:       v.Image,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
			DeletedAt:   v.DeletedAt,
		})
	}

	return dto.Response{
		Code:    200,
		Message: "APPROVE",
		Data:    productsResponse,
	}
}

func (p productService) UpdateTheProduct(ctx context.Context, req dto.ProductRequest) dto.Response {
	productUpdate, err := p.productRepository.FindByID(ctx, req.ID)
	if err != nil {
		return dto.Response{
			Code:    400,
			Message: "INVALID",
			Error:   err.Error(),
		}
	}

	if req == (dto.ProductRequest{}) {
		return dto.Response{
			Code:    400,
			Message: "INVALID",
			Error:   "body must bee fill",
		}
	}
	productUpdate.ID = req.ID
	productUpdate.Title = req.Title
	productUpdate.Description = req.Description
	productUpdate.Rating = req.Rating
	productUpdate.Image = req.Image
	productUpdate.UpdatedAt = time.Now()

	if err := p.productRepository.Update(ctx, &productUpdate); err != nil {
		return dto.Response{
			Code:    400,
			Message: "INVALID",
			Error:   err.Error(),
		}
	}

	return dto.Response{
		Code:    200,
		Message: "APPROVE",
	}
}
