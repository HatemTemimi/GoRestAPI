package product

import (
	models "apigo/internal/product/models"
	repository "apigo/internal/product/repository"
)

type ProductService struct {
	ProductRepository repository.ProductRepository
}

func ProvideProductService(p repository.ProductRepository) ProductService {
	return ProductService{ProductRepository: p}
}

func (p *ProductService) FindAll() []models.Product {
	return p.ProductRepository.FindAll()
}

func (p *ProductService) FindByID(id uint) models.Product {
	return p.ProductRepository.FindByID(id)
}

func (p *ProductService) Save(product models.Product) models.Product {
	p.ProductRepository.Save(product)

	return product
}

func (p *ProductService) Delete(product models.Product) {
	p.ProductRepository.Delete(product)
}
