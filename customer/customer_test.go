package customer

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCustomer(t *testing.T) {
	cust, err := NewCustomer("John", "Smith", "111-111-1111")
	assert.NotNil(t, cust)
	assert.Nil(t, err)

	cust, err = NewCustomer("John", "Smith", "(111)111-1111")
	assert.NotNil(t, cust)
	assert.Nil(t, err)

	cust, err = NewCustomer("John", "Smith", "(111) 111-1111")
	assert.NotNil(t, cust)
	assert.Nil(t, err)

	cust, err = NewCustomer("John", "Smith", "111 111 1111")
	assert.NotNil(t, cust)
	assert.Nil(t, err)

	cust, err = NewCustomer("John", "Smith", "111.111.1111")
	assert.NotNil(t, cust)
	assert.Nil(t, err)
}

func TestNewCustomerInvalidFirstName(t *testing.T) {
	cust, err := NewCustomer("", "Smith", "111-111-1111")
	assert.Nil(t, cust)
	assert.Error(t, err)
	assert.Equal(t, errors.New("Invalid first name, last name or phone number"), err)
}

func TestNewCustomerInvalidLastName(t *testing.T) {
	cust, err := NewCustomer("John", "", "111-111-1111")
	assert.Nil(t, cust)
	assert.Error(t, err)
	assert.Equal(t, errors.New("Invalid first name, last name or phone number"), err)
}

func TestNewCustomerInvalidPhoneNumber(t *testing.T) {
	cust, err := NewCustomer("John", "Smith", "111-111-111")
	assert.Nil(t, cust)
	assert.Error(t, err)
	assert.Equal(t, errors.New("Invalid first name, last name or phone number"), err)

	cust, err = NewCustomer("John", "Smith", "(111) 11 1111")
	assert.Nil(t, cust)
	assert.Error(t, err)
	assert.Equal(t, errors.New("Invalid first name, last name or phone number"), err)

	cust, err = NewCustomer("John", "Smith", "11111111111")
	assert.Nil(t, cust)
	assert.Error(t, err)
	assert.Equal(t, errors.New("Invalid first name, last name or phone number"), err)

	cust, err = NewCustomer("John", "Smith", "1111111111a")
	assert.Nil(t, cust)
	assert.Error(t, err)
	assert.Equal(t, errors.New("Invalid first name, last name or phone number"), err)

	cust, err = NewCustomer("John", "Smith", "((111)) 111-1111")
	assert.Nil(t, cust)
	assert.Error(t, err)
	assert.Equal(t, errors.New("Invalid first name, last name or phone number"), err)

	cust, err = NewCustomer("John", "Smith", "(111)-(111)-1111")
	assert.Nil(t, cust)
	assert.Error(t, err)
	assert.Equal(t, errors.New("Invalid first name, last name or phone number"), err)

	cust, err = NewCustomer("John", "Smith", "(111.111.1111)")
	assert.Nil(t, cust)
	assert.Error(t, err)
	assert.Equal(t, errors.New("Invalid first name, last name or phone number"), err)
}
