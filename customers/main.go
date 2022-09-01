package main

import (
	"customers/handlers"
	"flag"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	PORT string
)

func main() {

	flag.StringVar(&PORT, "port", "50090", "--port=50090")
	flag.Parse()

	if os.Getenv("PORT") != "" {
		PORT = os.Getenv("PORT")
	}

	r := gin.Default()

	r.GET("/greet", handlers.Greet)
	r.GET("/health", handlers.Health("ok"))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	//

	r.Run(":" + PORT) // http.ListenAndServe()

}
