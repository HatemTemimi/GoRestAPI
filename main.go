package main

import (
	"github.com/gin-gonic/gin"

	repository "apigo/internal/product/repository"

	service "apigo/internal/product/service"

	handler "apigo/internal/product/handlers"

	configs "apigo/config"
)

func main() {

	configs.InitEnvConfigs()
	var db configs.Database
	db.Create()

	productRepository := repository.ProvideProductRepostiory(db.Db)
	productService := service.ProvideProductService(productRepository)
	productAPI := handler.ProvideProductAPI(productService)

	r := gin.Default()

	r.GET("/products", productAPI.FindAll)
	r.GET("/products/:id", productAPI.FindByID)
	r.POST("/products", productAPI.Create)
	r.PUT("/products/:id", productAPI.Update)
	r.DELETE("/products/:id", productAPI.Delete)

	r.Run()

}
