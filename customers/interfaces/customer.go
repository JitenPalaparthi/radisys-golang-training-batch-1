package interfaces

import "customers/models"

type ICustomer interface {
	Add(*models.Customer) (*models.Customer, error)
	// Read(id string) *models.Customer
	// Update(id string, data map[string]interface{}) (*models.Customer, error)
	// Delete(id string) (int, error)
	// ReadAll() []models.Customer
}
