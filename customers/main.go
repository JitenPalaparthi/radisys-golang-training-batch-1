package main

import (
	"customers/dal"
	"customers/handlers"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	PORT string
	DSN  string
	//dsn := "host=localhost user=admin password=admin123 dbname=customersdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"

)

func main() {

	flag.StringVar(&PORT, "port", "50090", "--port=50090")
	flag.StringVar(&DSN, "db", "host=localhost user=admin password=admin123 dbname=customersdb port=5432 sslmode=disable TimeZone=Asia/Shanghai", "--db=host=localhost user=admin password=admin123 dbname=customersdb port=5432 sslmode=disable TimeZone=Asia/Shanghai")
	flag.Parse()

	if os.Getenv("PORT") != "" {
		PORT = os.Getenv("PORT")
	}
	if os.Getenv("DB_CONN") != "" {
		DSN = os.Getenv("DB_CONN")
	}

	db, err := dal.GetConnection(DSN)
	if err != nil {
		log.Fatal(err)
	}

	cdb := new(dal.CustomerDB)
	cdb.DB = db

	r := gin.Default()

	r.GET("/greet", handlers.Greet)
	r.GET("/health", handlers.Health("ok"))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	chandler := new(handlers.CustomerHandler)
	chandler.ICustomer = cdb
	v1 := r.Group("v1/public/customer")
	v1.POST("/add", chandler.Create())
	v1.GET("/:id", chandler.GetBy())
	v1.PUT("/:id", chandler.UpdateBy())
	v1.DELETE("/:id", chandler.DeleteBy())
	v1.GET("/all", chandler.GetAll())
	r.Run(":" + PORT) // http.ListenAndServe()
}
