package main

import (
	"apigo/product"

	configs "apigo/config"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("into main.go")

	configs.InitEnvConfigs()
	fmt.Println(configs.EnvConfigs)

	var db configs.Database
	db.Create()
	/*
		//fmt.Println(viper.Get("DB_USER"))
		dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=postgres port=%d sslmode=disable", configs.EnvConfigs.DB_USER, configs.EnvConfigs.DB_PASS, configs.EnvConfigs.DB_PORT)

		db, err := gorm.Open(postgres.New(postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true, // disables implicit prepared statement usage
		}), &gorm.Config{})
		if err != nil {
			panic(err)
		}
	*/

	productRepository := product.ProvideProductRepostiory(db.Db)
	productService := product.ProvideProductService(productRepository)
	productAPI := product.ProvideProductAPI(productService)

	r := gin.Default()

	r.GET("/products", productAPI.FindAll)
	r.GET("/products/:id", productAPI.FindByID)
	r.POST("/products", productAPI.Create)
	r.PUT("/products/:id", productAPI.Update)
	r.DELETE("/products/:id", productAPI.Delete)

	r.Run()

}
