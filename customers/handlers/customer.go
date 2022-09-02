package handlers

import (
	"customers/interfaces"
	"customers/messagebroker"
	"customers/models"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

type CustomerHandler struct {
	ICustomer interfaces.ICustomer
	Conn      []string
	// Logger    log.Logger
}

func (ch *CustomerHandler) Create() func(*gin.Context) {
	return func(c *gin.Context) {
		customer := new(models.Customer)
		err := c.Bind(customer)
		// var buf []byte
		// c.Request.Body.Read(buf)

		if err != nil {
			glog.Error(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
			c.Abort()
			return
		}

		if err := customer.Validate(); err != nil {
			glog.Error(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
			c.Abort()
			return
		}

		customer.Status = "created"
		customer.LastModified = fmt.Sprint(time.Now().Unix())
		glog.Info("creating a new customer")
		if cust, err := ch.ICustomer.Add(customer); err != nil {
			glog.Error(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
			c.Abort()
			return
		} else {
			// remove this code and write in examples/publish.go
			val, _ := customer.ToBytes()
			messagebroker.Publish(ch.Conn, "customer.created", []byte(strconv.Itoa(customer.ID)), val)
			c.JSON(http.StatusCreated, cust)
			return
		}

	}
}

func (ch *CustomerHandler) GetBy() func(*gin.Context) {
	return func(c *gin.Context) {
		_id, ok := c.Params.Get("id")
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "invalid id"})
			c.Abort()
			return
		}

		id, err := strconv.Atoi(_id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
			c.Abort()
			return
		}
		customer, err := ch.ICustomer.GetBy(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, customer)
		c.Abort()
	}
}

func (ch *CustomerHandler) UpdateBy() func(*gin.Context) {
	return func(c *gin.Context) {
		_id, ok := c.Params.Get("id")
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "invalid id"})
			c.Abort()
			return
		}

		id, err := strconv.Atoi(_id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
			c.Abort()
			return
		}

		data := make(map[string]interface{})

		err = c.Bind(&data)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
			c.Abort()
			return
		}

		customer, err := ch.ICustomer.UpdateBy(id, data)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, customer)
		c.Abort()
	}
}

func (ch *CustomerHandler) DeleteBy() func(*gin.Context) {
	return func(c *gin.Context) {
		_id, ok := c.Params.Get("id")
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "invalid id"})
			c.Abort()
			return
		}

		id, err := strconv.Atoi(_id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
			c.Abort()
			return
		}

		n, err := ch.ICustomer.DeleteBy(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, n)
		c.Abort()
	}
}

func (ch *CustomerHandler) GetAll() func(*gin.Context) {
	return func(c *gin.Context) {
		customers, err := ch.ICustomer.GetAll()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, customers)
		c.Abort()
	}
}
