package scheduler

import (
	"service_scheduler/customer"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckIn(t *testing.T) {
	s := NewScheduler()

	// Customer check-ins
	customer1 := &customer.Customer{FirstName: "John", LastName: "Smith", PhoneNumber: "123-456-7890", IsVIP: false}
	customer2 := &customer.Customer{FirstName: "Alice", LastName: "Johnson", PhoneNumber: "987-654-3210", IsVIP: false}
	customer3 := &customer.Customer{FirstName: "Bob", LastName: "Meyers", PhoneNumber: "456-987-9638", IsVIP: false}
	customer4 := &customer.Customer{FirstName: "Mike", LastName: "Tomlin", PhoneNumber: "567-890-8527", IsVIP: true}
	customer5 := &customer.Customer{FirstName: "Clyde", LastName: "Johnson", PhoneNumber: "245-753-0025", IsVIP: true}
	customer6 := &customer.Customer{FirstName: "Alex", LastName: "Jennings", PhoneNumber: "852-356-3468", IsVIP: false}
	customer7 := &customer.Customer{FirstName: "Jason", LastName: "Wright", PhoneNumber: "246-846-5132", IsVIP: true}
	customer8 := &customer.Customer{FirstName: "Carla", LastName: "Adams", PhoneNumber: "410-212-9987", IsVIP: true}
	customer9 := &customer.Customer{FirstName: "Gwen", LastName: "Tyler", PhoneNumber: "212-458-5867", IsVIP: false}
	customer0 := &customer.Customer{FirstName: "Virginia", LastName: "Jefferson", PhoneNumber: "918-361-1124", IsVIP: true}

	// Check-in customers
	s.CheckIn(customer1)
	s.CheckIn(customer2)
	s.CheckIn(customer3)
	s.CheckIn(customer4)
	s.CheckIn(customer5)
	s.CheckIn(customer6)
	s.CheckIn(customer7)
	s.CheckIn(customer8)
	s.CheckIn(customer9)
	s.CheckIn(customer0)

	assert.Equal(t, len(s.RegularQueue), 5)
	assert.Equal(t, len(s.VipQueue), 5)
}

func TestGetNextCustomer(t *testing.T) {
	s := NewScheduler()

	// Customer check-ins
	customer1 := &customer.Customer{FirstName: "John", LastName: "Smith", PhoneNumber: "123-456-7890", IsVIP: false}
	customer2 := &customer.Customer{FirstName: "Alice", LastName: "Johnson", PhoneNumber: "987-654-3210", IsVIP: false}
	customer3 := &customer.Customer{FirstName: "Bob", LastName: "Meyers", PhoneNumber: "456-987-9638", IsVIP: false}
	customer4 := &customer.Customer{FirstName: "Mike", LastName: "Tomlin", PhoneNumber: "567-890-8527", IsVIP: true}
	customer5 := &customer.Customer{FirstName: "Clyde", LastName: "Jones", PhoneNumber: "245-753-0025", IsVIP: true}
	customer6 := &customer.Customer{FirstName: "Alex", LastName: "Jennings", PhoneNumber: "852-356-3468", IsVIP: false}
	customer7 := &customer.Customer{FirstName: "Jason", LastName: "Wright", PhoneNumber: "246-846-5132", IsVIP: true}
	customer8 := &customer.Customer{FirstName: "Carla", LastName: "Adams", PhoneNumber: "410-212-9987", IsVIP: true}
	customer9 := &customer.Customer{FirstName: "Gwen", LastName: "Tyler", PhoneNumber: "212-458-5867", IsVIP: false}
	customer0 := &customer.Customer{FirstName: "Virginia", LastName: "Jefferson", PhoneNumber: "918-361-1124", IsVIP: true}

	// Check-in customers
	s.CheckIn(customer1)
	s.CheckIn(customer2)
	s.CheckIn(customer3)
	s.CheckIn(customer4)
	s.CheckIn(customer5)
	s.CheckIn(customer6)
	s.CheckIn(customer7)
	s.CheckIn(customer8)
	s.CheckIn(customer9)
	s.CheckIn(customer0)

	assert.Equal(t, len(s.RegularQueue), 5)
	assert.Equal(t, len(s.VipQueue), 5)

	cust := s.GetNextCustomer()
	assert.True(t, cust.IsVIP)
	assert.Equal(t, cust.FirstName+" "+cust.LastName, "Mike Tomlin")
	assert.Equal(t, len(s.VipQueue), 4)
	assert.Equal(t, len(s.RegularQueue), 5)
	assert.Equal(t, s.vipRate, 1)
}

func TestGetNextCustomerVIPProcessingRate(t *testing.T) {
	s := NewScheduler()

	// Customer check-ins
	customer1 := &customer.Customer{FirstName: "John", LastName: "Smith", PhoneNumber: "123-456-7890", IsVIP: false}
	customer2 := &customer.Customer{FirstName: "Alice", LastName: "Johnson", PhoneNumber: "987-654-3210", IsVIP: false}
	customer3 := &customer.Customer{FirstName: "Bob", LastName: "Meyers", PhoneNumber: "456-987-9638", IsVIP: false}
	customer4 := &customer.Customer{FirstName: "Mike", LastName: "Tomlin", PhoneNumber: "567-890-8527", IsVIP: true}
	customer5 := &customer.Customer{FirstName: "Clyde", LastName: "Jones", PhoneNumber: "245-753-0025", IsVIP: true}
	customer6 := &customer.Customer{FirstName: "Alex", LastName: "Jennings", PhoneNumber: "852-356-3468", IsVIP: false}
	customer7 := &customer.Customer{FirstName: "Jason", LastName: "Wright", PhoneNumber: "246-846-5132", IsVIP: true}
	customer8 := &customer.Customer{FirstName: "Carla", LastName: "Adams", PhoneNumber: "410-212-9987", IsVIP: true}
	customer9 := &customer.Customer{FirstName: "Gwen", LastName: "Tyler", PhoneNumber: "212-458-5867", IsVIP: false}
	customer0 := &customer.Customer{FirstName: "Virginia", LastName: "Jefferson", PhoneNumber: "918-361-1124", IsVIP: true}

	// Check-in customers
	s.CheckIn(customer1)
	s.CheckIn(customer2)
	s.CheckIn(customer3)
	s.CheckIn(customer4)
	s.CheckIn(customer5)
	s.CheckIn(customer6)
	s.CheckIn(customer7)
	s.CheckIn(customer8)
	s.CheckIn(customer9)
	s.CheckIn(customer0)

	assert.Equal(t, len(s.RegularQueue), 5)
	assert.Equal(t, len(s.VipQueue), 5)

	cust := s.GetNextCustomer()
	assert.True(t, cust.IsVIP)
	assert.Equal(t, cust.FirstName+" "+cust.LastName, "Mike Tomlin")
	assert.Equal(t, len(s.VipQueue), 4)
	assert.Equal(t, len(s.RegularQueue), 5)
	assert.Equal(t, s.vipRate, 1)

	cust = s.GetNextCustomer()
	assert.True(t, cust.IsVIP)
	assert.Equal(t, cust.FirstName+" "+cust.LastName, "Clyde Jones")
	assert.Equal(t, len(s.VipQueue), 3)
	assert.Equal(t, len(s.RegularQueue), 5)
	assert.Equal(t, s.vipRate, 0)

	cust = s.GetNextCustomer()
	assert.False(t, cust.IsVIP)
	assert.Equal(t, cust.FirstName+" "+cust.LastName, "John Smith")
	assert.Equal(t, len(s.VipQueue), 3)
	assert.Equal(t, len(s.RegularQueue), 4)
	assert.Equal(t, s.vipRate, 2)

	cust = s.GetNextCustomer()
	assert.True(t, cust.IsVIP, true)
	assert.Equal(t, cust.FirstName+" "+cust.LastName, "Jason Wright")
	assert.Equal(t, len(s.VipQueue), 2)
	assert.Equal(t, len(s.RegularQueue), 4)
	assert.Equal(t, s.vipRate, 1)

}

func TestGetNextCustomerVIPProcessingRateWhenThereIsNoRegularCust(t *testing.T) {
	s := NewScheduler()

	// Customer check-ins
	customer4 := &customer.Customer{FirstName: "Mike", LastName: "Tomlin", PhoneNumber: "567-890-8527", IsVIP: true}
	customer5 := &customer.Customer{FirstName: "Clyde", LastName: "Jones", PhoneNumber: "245-753-0025", IsVIP: true}
	customer7 := &customer.Customer{FirstName: "Jason", LastName: "Wright", PhoneNumber: "246-846-5132", IsVIP: true}
	customer8 := &customer.Customer{FirstName: "Carla", LastName: "Adams", PhoneNumber: "410-212-9987", IsVIP: true}
	customer0 := &customer.Customer{FirstName: "Virginia", LastName: "Jefferson", PhoneNumber: "918-361-1124", IsVIP: true}

	// Check-in customers
	s.CheckIn(customer4)
	s.CheckIn(customer5)
	s.CheckIn(customer7)
	s.CheckIn(customer8)
	s.CheckIn(customer0)

	assert.Equal(t, len(s.RegularQueue), 0)
	assert.Equal(t, len(s.VipQueue), 5)

	cust := s.GetNextCustomer()
	assert.True(t, cust.IsVIP)
	assert.Equal(t, cust.FirstName+" "+cust.LastName, "Mike Tomlin")
	assert.Equal(t, len(s.VipQueue), 4)
	assert.Equal(t, len(s.RegularQueue), 0)
	assert.Equal(t, s.vipRate, 1)

	cust = s.GetNextCustomer()
	assert.True(t, cust.IsVIP)
	assert.Equal(t, cust.FirstName+" "+cust.LastName, "Clyde Jones")
	assert.Equal(t, len(s.VipQueue), 3)
	assert.Equal(t, len(s.RegularQueue), 0)
	assert.Equal(t, s.vipRate, 0)

	cust = s.GetNextCustomer()
	assert.True(t, cust.IsVIP)
	assert.Equal(t, cust.FirstName+" "+cust.LastName, "Jason Wright")
	assert.Equal(t, len(s.VipQueue), 2)
	assert.Equal(t, len(s.RegularQueue), 0)
	assert.Equal(t, s.vipRate, 0)

	cust = s.GetNextCustomer()
	assert.True(t, cust.IsVIP, true)
	assert.Equal(t, cust.FirstName+" "+cust.LastName, "Carla Adams")
	assert.Equal(t, len(s.VipQueue), 1)
	assert.Equal(t, len(s.RegularQueue), 0)
	assert.Equal(t, s.vipRate, 0)

	cust = s.GetNextCustomer()
	assert.True(t, cust.IsVIP, true)
	assert.Equal(t, cust.FirstName+" "+cust.LastName, "Virginia Jefferson")
	assert.Equal(t, len(s.VipQueue), 0)
	assert.Equal(t, len(s.RegularQueue), 0)
	assert.Equal(t, s.vipRate, 0)

}
