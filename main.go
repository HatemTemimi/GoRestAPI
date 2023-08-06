package main

import (
	"apigo/product"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=localhost user=SuperUser password=SuperSecure dbname=postgres port=5432 sslmode=disable",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	productRepository := product.ProvideProductRepostiory(db)
	productService := product.ProvideProductService(productRepository)
	productAPI := product.ProvideProductAPI(productService)

	db.AutoMigrate(&product.Product{})

	r := gin.Default()

	r.GET("/products", productAPI.FindAll)
	r.GET("/products/:id", productAPI.FindByID)
	r.POST("/products", productAPI.Create)
	r.PUT("/products/:id", productAPI.Update)
	r.DELETE("/products/:id", productAPI.Delete)

	r.Run()

}
