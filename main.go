package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var products = []product{
	{ID: 1, Name: "Car", Price: 90000.00},
	{ID: 2, Name: "Tyre", Price: 560.00},
	{ID: 3, Name: "Tipper motor", Price: 56700.00},
}

func main() {
	r := gin.Default()
	r.GET("/products", getProduct)
	r.POST("/products", postProduct)
	r.GET("/products/:id", getProductByID)
	r.PATCH("/products/:id", patchProduct)
	r.DELETE("/products/:id", deleteProduct)
	r.Run(":8080")
}

func getProduct(c *gin.Context) {
	c.JSON(http.StatusOK, products)
}

func postProduct(c *gin.Context) {
	var newProduct product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newProduct.ID = len(products) + 1
	products = append(products, newProduct)
	c.JSON(http.StatusOK, gin.H{"message": "product posted successfully"})
}

func patchProduct(c *gin.Context) {
	var patchProduct product
	if err := c.ShouldBindJSON(&patchProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}

	for i, p := range products {
		if p.ID == id {
			if patchProduct.Name != "" {
				products[i].Name = patchProduct.Name
			}
			if patchProduct.Price != 0 {
				products[i].Price = patchProduct.Price
			}
			c.JSON(http.StatusOK, products[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
}

func deleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}

	for i, p := range products {
		if p.ID == id {
			products = append(products[:i], products[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "product deleted successfully"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
}

func getProductByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}

	for _, p := range products {
		if p.ID == id {
			c.JSON(http.StatusOK, p)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
}
