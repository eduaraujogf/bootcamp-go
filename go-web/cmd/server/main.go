package main

import (
	"log"

	"github.com/eduaraujogf/bootcamp-go/go-web/cmd/server/controllers"
	"github.com/eduaraujogf/bootcamp-go/go-web/internal/email"
	"github.com/eduaraujogf/bootcamp-go/go-web/internal/products"
	store "github.com/eduaraujogf/bootcamp-go/go-web/pkg"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("erro ao carregar o arquivo .env")
	}

	db := store.New(store.FileType, "go-web/products.json")
	//1. repositório
	repo := products.NewRepository(db)

	//2. serviço (regra de negócio)
	//emailSES := email.NewSES()
	emailSendGrid := email.NewSendgrid()
	service := products.NewService(repo, emailSendGrid)

	//3. controller
	p := controllers.NewProduct(service)

	r := gin.Default()
	pr := r.Group("/products")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.PUT("/:id", p.Update())
	pr.PATCH("/:id", p.UpdateName())
	pr.DELETE("/:id", p.Delete())
	r.Run()
}
