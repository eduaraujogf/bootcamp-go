package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func rootHandler(c *gin.Context) {
	//Query string (url?name=""&lastname="")
	name := c.Query("name")
	lastName := c.Query("lastname")
	log.Println(name, lastName)

	c.JSON(http.StatusOK, gin.H{
		"message": "queries",
	})

	// //The body, header and method are inside of the gin context.
	// body := c.Request.Body
	// header := c.Request.Header
	// method := c.Request.Method

	// fmt.Println("Eu recebi algo!")
	// fmt.Printf("\tMétodo: %s\n", method)
	// fmt.Printf("\tConteúdo do cabeçalho:\n")

	// for key, value := range header {
	// 	fmt.Printf("\t\t%s -> %s\n", key, value)
	// }

	// fmt.Printf("\tO body é um io.ReadCloser:(%s), e para trabalhar com ele teremos que ler depois\n", body)
	// fmt.Println("¡Yay!")
	// c.String(200, "Eu recebi!") //Respondemos ao cliente com 200 OK e uma mensagem.
}

func helloHandler(c *gin.Context) {
	// URL Param (url/:id)
	id := c.Param("id")

	_, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "id is not a number",
		})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Hello: " + id,
	})
}

type Employee struct {
}

func employeeHandler(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{})
}

func main() {
	gin.SetMode("debug")

	router := gin.Default()

	router.GET("/", rootHandler)

	group := router.Group("/api/v1")
	{
		group.GET("/hello-world/:id", helloHandler)
		group.POST("/employee", employeeHandler)
	}

	router.Run()
}
