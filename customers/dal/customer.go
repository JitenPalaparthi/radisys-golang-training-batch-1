package dal

import (
	"customers/models"
	"fmt"

	"gorm.io/gorm"
)

type CustomerDB struct {
	Client *gorm.DB // not a promoted field so everytime call Client.Func
	//*gorm.DB          // promoted filed
}

func (c *CustomerDB) Add(customer *models.Customer) (*models.Customer, error) {
	//c.Client.AutoMigrate(&models.Customer{})
	c.Client.AutoMigrate(&models.Customer{})
	tx := c.Client.Create(customer)
	if tx.Error != nil {
		fmt.Println("error here?---")
		return nil, tx.Error
	}
	return customer, nil
}
