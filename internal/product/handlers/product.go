package product

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	models "apigo/internal/product/models"
	service "apigo/internal/product/service"

	_ "github.com/swaggo/swag/example/celler/httputil"
)

type ProductAPI struct {
	ProductService service.ProductService
}

func ProvideProductAPI(p service.ProductService) ProductAPI {
	return ProductAPI{ProductService: p}
}

// FindAll finds all products
//
//	@Summary		List products
//	@Description	get products
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.Product
//	@Failure		400	{object}	error
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/products [get]
func (p *ProductAPI) FindAll(c *gin.Context) {
	products := p.ProductService.FindAll()
	c.JSON(http.StatusOK, gin.H{
		"message":  "All Products",
		"products": models.ToProductDTOs(products),
	})
}

// FindByID finds one product by ID
//
//	@Summary		finds one product by ID
//	@Description	gets product
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}		models.Product
//	@Failure		400	{object}	error
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/products/{id} [get]
func (p *ProductAPI) FindByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product := p.ProductService.FindByID(uint(id))

	c.JSON(http.StatusOK, gin.H{"product": models.ToProductDTO(product)})
}

// Create	creates a product
//
//	@Summary		creates a product
//	@Description	creates a product
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}		models.Product
//	@Failure		400	{object}	error
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/products [post]
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

// update	updates a product
//
//	@Summary		updates a product
//	@Description	updates a product
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}		models.Product
//	@Failure		400	{object}	error
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/products/{id} [post]
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

// update	deletes a product
//
//	@Summary		deletes a product
//	@Description	deletes a product
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}		models.Product
//	@Failure		400	{object}	error
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/products/{id} [delete]
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
