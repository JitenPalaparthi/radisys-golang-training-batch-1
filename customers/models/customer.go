package models

import (
	"encoding/json"
	"errors"
)

var (
	ErrInvalidName      = errors.New("invalid name")
	ErrInvalidEmail     = errors.New("invalid email")
	ErrInvalidContactNo = errors.New("invalid contactNo")
)

type Customer struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Address      string `json:"address"`
	ContactNo    string `json:"contactNo" gorm:"column:contactNo"`
	Status       string `json:"status"`
	LastModified string `json:"lastModified" gorm:"column:lastModified"`
}

func (c *Customer) Validate() error {
	if c.Name == "" {
		return ErrInvalidName
	}
	if c.Email == "" {
		return ErrInvalidEmail
	}
	if c.ContactNo == "" {
		return ErrInvalidContactNo
	}
	return nil
}

func (c *Customer) ToString() (string, error) {
	s, err := c.ToBytes()
	return string(s), err
}

func (c *Customer) ToBytes() ([]byte, error) {
	if buf, err := json.Marshal(c); err != nil {
		return nil, err
	} else {
		return buf, nil
	}
}
