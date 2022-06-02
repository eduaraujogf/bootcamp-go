package main

import (
	"log"
	"net/http"
	"os"

	docs "github.com/eduaraujogf/bootcamp-go/docs"
	"github.com/eduaraujogf/bootcamp-go/go-web/cmd/server/controllers"
	"github.com/eduaraujogf/bootcamp-go/go-web/internal/email"
	"github.com/eduaraujogf/bootcamp-go/go-web/internal/products"
	"github.com/eduaraujogf/bootcamp-go/go-web/pkg/store"
	"github.com/eduaraujogf/bootcamp-go/go-web/pkg/web"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")

	// We want to make sure the token is set, bail if not
	if requiredToken == "" {
		log.Fatal("Please set token environment variable")
	}

	return func(c *gin.Context) {
		token := c.GetHeader("token")

		if token == "" {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				web.DecodeError(http.StatusUnauthorized, "Token vazio"),
			)
			return
		}

		if token != requiredToken {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				web.DecodeError(http.StatusUnauthorized, "Token inválido"),
			)
			return
		}

		c.Next()
	}
}

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

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

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	pr := r.Group("/products")
	{
		pr.Use(TokenAuthMiddleware())

		pr.POST("/", p.Store())
		pr.GET("/", p.GetAll())
		pr.PUT("/:id", p.Update())
		pr.PATCH("/:id", p.UpdateName())
		pr.DELETE("/:id", p.Delete())
	}

	r.Run()
}
