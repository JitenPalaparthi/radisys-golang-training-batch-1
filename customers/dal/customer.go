package dal

import (
	"customers/models"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type CustomerDB struct {
	//Client *gorm.DB // not a promoted field so everytime call Client.Func
	*gorm.DB // promoted filed
}

func (c *CustomerDB) Add(customer *models.Customer) (*models.Customer, error) {
	//c.Client.AutoMigrate(&models.Customer{})
	c.AutoMigrate(&models.Customer{})
	tx := c.Create(customer)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return customer, nil
}

func (c *CustomerDB) GetBy(id int) (*models.Customer, error) {
	customer := new(models.Customer)
	tx := c.First(customer, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return customer, nil
}

func (c *CustomerDB) GetByStatus(status string) ([]models.Customer, error) {
	var customers []models.Customer
	tx := c.Find(&customers, map[string]interface{}{"status": status})
	if tx.Error != nil {
		return nil, tx.Error
	}
	return customers, nil
}

func (c *CustomerDB) UpdateBy(id int, data map[string]interface{}) (*models.Customer, error) {
	customer := new(models.Customer)
	customer.ID = id
	data["lastModified"] = fmt.Sprint(time.Now().Unix())
	tx := c.Model(customer).Updates(data)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return c.GetBy(id)
}

func (c *CustomerDB) DeleteBy(id int) (int, error) {
	tx := c.Delete(&models.Customer{}, id)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (c *CustomerDB) GetAll() ([]models.Customer, error) {
	var customers []models.Customer
	tx := c.Find(&customers)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return customers, nil
}
