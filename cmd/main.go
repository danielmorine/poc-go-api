package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()

	if err != nil {
		panic(err)
	}

	ProductRepository := repository.NewProductRepository(dbConnection)

	ProductUseCase := usecase.NewProductUseCase(ProductRepository)

	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/products", ProductController.GetProduts)
	server.GET("/product/:productId", ProductController.GetProductById)
	server.POST("/product", ProductController.CreateProduct)
	server.PUT("/product", ProductController.UpdateProduct)
	server.DELETE("/product/:productId", ProductController.DeleteProduct)

	server.Run(":8000")
}
