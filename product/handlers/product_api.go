package product

import (
	models "apigo/product/models"
	service "apigo/product/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductAPI struct {
	ProductService service.ProductService
}

func ProvideProductAPI(p service.ProductService) ProductAPI {
	return ProductAPI{ProductService: p}
}

func (p *ProductAPI) FindAll(c *gin.Context) {
	products := p.ProductService.FindAll()
	c.JSON(http.StatusOK, gin.H{
		"message":  "All Products",
		"products": models.ToProductDTOs(products),
	})
}

func (p *ProductAPI) FindByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product := p.ProductService.FindByID(uint(id))

	c.JSON(http.StatusOK, gin.H{"product": models.ToProductDTO(product)})
}

func (p *ProductAPI) Create(c *gin.Context) {
	var productDTO models.ProductDTO
	err := c.BindJSON(&productDTO)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	createdProduct := p.ProductService.Save(models.ToProduct(productDTO))

	c.JSON(http.StatusOK, gin.H{
		"message": "Product added",
		"product": models.ToProductDTO(createdProduct)})
}

func (p *ProductAPI) Update(c *gin.Context) {
	var productDTO models.ProductDTO
	err := c.BindJSON(&productDTO)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	product := p.ProductService.FindByID(uint(id))
	if product == (models.Product{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	product.Code = productDTO.Code
	product.Price = productDTO.Price
	p.ProductService.Save(product)

	c.Status(http.StatusOK)
}

func (p *ProductAPI) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product := p.ProductService.FindByID(uint(id))
	if product == (models.Product{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	p.ProductService.Delete(product)

	c.Status(http.StatusOK)
}
