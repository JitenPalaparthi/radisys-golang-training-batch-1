package handlers

import (
	"customers/interfaces"
	"customers/models"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	ICustomer interfaces.ICustomer
	Logger    log.Logger
}

func (ch *CustomerHandler) Create() func(*gin.Context) {
	return func(c *gin.Context) {
		customer := new(models.Customer)
		err := c.Bind(customer)
		// var buf []byte
		// c.Request.Body.Read(buf)

		if err != nil {
			//ch.Logger.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"status": "fail-1", "message": err})
			c.Abort()
			return
		}

		if err := customer.Validate(); err != nil {
			//ch.Logger.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"status": "fail-2", "message": err})
			c.Abort()
			return
		}

		customer.Status = "created"
		customer.LastModified = fmt.Sprint(time.Now().Unix())
		if cust, err := ch.ICustomer.Add(customer); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "fail-3", "message": err})
			c.Abort()
			return
		} else {
			c.JSON(http.StatusCreated, cust)
			return
		}

	}
}
