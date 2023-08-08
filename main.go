package main

import (
	"github.com/gin-gonic/gin"

	repository "apigo/internal/product/repository"

	service "apigo/internal/product/service"

	handler "apigo/internal/product/handlers"

	userHandler "apigo/internal/user/handlers"

	middleware "apigo/internal/middlewares"

	configs "apigo/config"

	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	swaggerFiles "github.com/swaggo/files"

	_ "apigo/docs"
)

//	@title			Products API
//	@version		1.0
//	@description	REST API implementation with Go

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/products

func main() {

	configs.InitEnvConfigs()
	var db configs.Database
	db.Setup()

	productRepository := repository.ProvideProductRepostiory(db.Db)
	productService := service.ProvideProductService(productRepository)
	productAPI := handler.ProvideProductAPI(productService)

	userAPI := userHandler.MakeUserApi(db.Db)

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/products", middleware.RequireAuth, productAPI.FindAll)
	r.GET("/products/:id", productAPI.FindByID)
	r.POST("/products", productAPI.Create)
	r.PUT("/products/:id", productAPI.Update)
	r.DELETE("/products/:id", productAPI.Delete)

	r.POST("/auth/signup", userAPI.Signup)
	r.POST("/auth/login", userAPI.Login)

	r.Run()

}
