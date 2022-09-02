package interfaces

import "customers/models"

type ICustomer interface {
	Add(*models.Customer) (*models.Customer, error)
	GetBy(id int) (*models.Customer, error)
	UpdateBy(id int, data map[string]interface{}) (*models.Customer, error)
	DeleteBy(id int) (int, error)
	GetAll() ([]models.Customer, error)
}
