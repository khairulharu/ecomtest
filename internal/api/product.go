package api

import (
	"ecomtest/domain"
	"ecomtest/dto"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type productApi struct {
	productService domain.ProductService
}

func NewProduct(productService domain.ProductService, app *fiber.App) {
	h := productApi{
		productService: productService,
	}

	app.Get("/products", h.ShowProducts)
	app.Get("/product", h.GetDetailOfProduct)
	app.Post("/product", h.CreateNewProduct)
	app.Put("/product", h.UpdateTheProduct)
	app.Delete("product", h.DeleteTheProduct)
}

func (p productApi) ShowProducts(ctx *fiber.Ctx) error {
	res := p.productService.ShowAllProducts(ctx.Context())
	return ctx.Status(int(res.Code)).JSON(res)
}

func (p productApi) GetDetailOfProduct(ctx *fiber.Ctx) error {
	idString := ctx.Query("id")
	id, _ := strconv.Atoi(idString)
	res := p.productService.GetDetailProduct(ctx.Context(), int64(id))
	return ctx.Status(int(res.Code)).JSON(res)
}

func (p productApi) CreateNewProduct(ctx *fiber.Ctx) error {
	var reqProduct dto.ProductRequest
	if err := ctx.BodyParser(&reqProduct); err != nil {
		return ctx.SendStatus(400)
	}

	res := p.productService.CreateNewProduct(ctx.Context(), reqProduct)
	return ctx.Status(int(res.Code)).JSON(res)
}

func (p productApi) UpdateTheProduct(ctx *fiber.Ctx) error {
	var reqProduct dto.ProductRequest
	idString := ctx.Query("id")
	id, _ := strconv.Atoi(idString)
	reqProduct.ID = int64(id)
	if err := ctx.BodyParser(&reqProduct); err != nil {
		return ctx.SendStatus(400)
	}

	res := p.productService.UpdateTheProduct(ctx.Context(), reqProduct)

	return ctx.Status(int(res.Code)).JSON(res)
}

func (p productApi) DeleteTheProduct(ctx *fiber.Ctx) error {
	idString := ctx.Query("id")
	id, _ := strconv.Atoi(idString)
	res := p.productService.DeleteTheProduct(ctx.Context(), int64(id))
	return ctx.Status(int(res.Code)).JSON(res)
}
