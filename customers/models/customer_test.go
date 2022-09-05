package models_test

import (
	"customers/models"
	"testing"
)

func TestValidatePositive(t *testing.T) {
	customer := new(models.Customer)
	customer.Name = "Customer-1"
	customer.Email = "Customer@Customer.Com"
	customer.ContactNo = "111111111"

	err := customer.Validate()

	if err != nil {
		t.Log("customer must not return err.")
		t.Fail()
	}
}

func TestValidateNegative(t *testing.T) {
	customer := new(models.Customer)
	customer.Name = "Customer-1"
	customer.Email = "Customer@Customer.Com"
	//customer.ContactNo = "111111111"

	err := customer.Validate()

	if err == nil {
		t.Log("customer must return err.")
		t.Fail()
	}
}

func TestToString(t *testing.T) {
	customer := new(models.Customer)
	customer.Name = "Customer-1"
	customer.Email = "Customer@Customer.Com"
	customer.ContactNo = "111111111"
	actualResult, _ := customer.ToString()
	expectedResult := `{"id":0,"name":"Customer-1","email":"Customer@Customer.Com","address":"","contactNo":"111111111","status":"","lastModified":""}`

	if actualResult != expectedResult {
		t.Log(actualResult)
		t.Fail()
	}
}
