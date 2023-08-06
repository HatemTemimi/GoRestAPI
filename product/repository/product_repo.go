package repository

import (
	model "apigo/product/models"
	"fmt"

	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func ProvideProductRepostiory(DB *gorm.DB) ProductRepository {
	return ProductRepository{DB: DB}
}

func (p *ProductRepository) FindAll() []model.Product {
	var products []model.Product
	if p.DB.Find(&products).Error != nil {
		fmt.Println("labes")
	} else {
		fmt.Println("mochkla")
	}
	return products
}

func (p *ProductRepository) FindByID(id uint) model.Product {
	var product model.Product
	p.DB.First(&product, id)

	return product
}

func (p *ProductRepository) Save(product model.Product) model.Product {
	p.DB.Save(&product)

	return product
}

func (p *ProductRepository) Delete(product model.Product) {
	p.DB.Delete(&product)
}
