package main

import (
	"customers/dal"
	"customers/handlers"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

var (
	PORT string
	DSN  string
	//dsn := "host=localhost user=admin password=admin123 dbname=customersdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	KAFKA_CONN []string = []string{"localhost:29092"}
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: example -stderrthreshold=[INFO|WARNING|ERROR|FATAL] -log_dir=[string]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func init() {
	flag.Usage = usage
	// NOTE: This next line is key you have to call flag.Parse() for the command line
	// options or "flags" that are defined in the glog module to be picked up.
	//flag.Parse()
}

func main() {

	flag.StringVar(&PORT, "port", "50090", "--port=50090")
	flag.StringVar(&DSN, "db", "host=localhost user=admin password=admin123 dbname=customersdb port=5432 sslmode=disable TimeZone=Asia/Shanghai", "--db=host=localhost user=admin password=admin123 dbname=customersdb port=5432 sslmode=disable TimeZone=Asia/Shanghai")

	flag.Set("stderrthreshold", "INFO") // can set up the glog
	flag.Parse()
	defer glog.Flush()
	if os.Getenv("PORT") != "" {
		PORT = os.Getenv("PORT")
	}
	if os.Getenv("DB_CONN") != "" {
		DSN = os.Getenv("DB_CONN")
	}

	glog.Info("Connecting to the database--")
	db, err := dal.GetConnection(DSN)

	if err != nil {
		glog.Fatal(err)
	}

	glog.Info("creating new instance of CustomerDB")
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
	chandler.Conn = KAFKA_CONN

	v1 := r.Group("v1/public/customer")
	v1.POST("/add", chandler.Create())
	v1.GET("/:id", chandler.GetBy())
	v1.PUT("/:id", chandler.UpdateBy())
	v1.DELETE("/:id", chandler.DeleteBy())
	v1.GET("/all", chandler.GetAll())
	r.Run(":" + PORT) // http.ListenAndServe()
}

// observability
// logs
// metrics
// traces
//
