package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Id       int64   `json:"id"`
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

type Products []Product

var lastID int64
var p Products

func getAll(c *gin.Context) {
	c.JSON(http.StatusOK, &p)
}

func products(c *gin.Context) {
	token := c.GetHeader("token")

	if token != "123456" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid token",
		})
		return
	}
	var product Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	lastID++
	product.Id = lastID

	p = append(p, product)

	c.JSON(200, gin.H{
		"data": product,
	})
}

func main() {
	gin.SetMode("debug")

	router := gin.Default()

	group := router.Group("/products")
	{
		group.POST("/", products)
		group.GET("/", getAll)
	}

	router.Run()
}
