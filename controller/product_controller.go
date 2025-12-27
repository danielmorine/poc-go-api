package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productUseCase usecase.ProductUseCase
}

func NewProductController(usecase usecase.ProductUseCase) productController {
	return productController{
		productUseCase: usecase,
	}
}

func (p *productController) GetProduts(ctx *gin.Context) {

	response, err := p.productUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (p *productController) CreateProduct(ctx *gin.Context) {
	var product model.Product
	err := ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Status: false, Message: err.Error(), Data: nil})
		return
	}

	insertedProduct, err := p.productUseCase.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, insertedProduct)
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}

func (p *productController) UpdateProduct(ctx *gin.Context) {
	var product model.Product
	err := ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Status: false, Message: err.Error(), Data: nil})
		return
	}

	response, err := p.productUseCase.UpdateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (p *productController) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("productId")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: "Id do produto não pode ser nulo",
			Data:    nil,
		})
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: "Id do produto precisa ser um número",
			Data:    nil,
		})
		return
	}

	response, err := p.productUseCase.DeleteProduct(productId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (p *productController) GetProductById(ctx *gin.Context) {
	id := ctx.Param("productId")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: "Id do produto não pode ser nulo",
			Data:    nil,
		})
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: "Id do produto precisa ser um número",
			Data:    nil,
		})
		return
	}

	response, err := p.productUseCase.GetProductById(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	if response.Data == nil {
		ctx.JSON(http.StatusNotFound, model.Response{
			Status:  true,
			Message: "Produto não encontrado",
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
