package scheduler

import (
	"fmt"
	"service_scheduler/customer"
	"sync"
)

const vipProcessingRate = 2
const ServiceTime = 5

type ServiceScheduler struct {
	RegularQueue   []*customer.Customer
	VipQueue       []*customer.Customer
	lock           sync.Mutex
	vipRate        int
	totalCheckedIn int
}

// Initializes the ServiceScheduler with a predefined `vipProcessingRate` and returns a pointer to the object.
func NewScheduler() *ServiceScheduler {
	return &ServiceScheduler{vipRate: vipProcessingRate}
}

// Checks in a customer into the appropriate queue based on whether they are a VIP or not while
// providing the longest time the customer could be waiting
func (s *ServiceScheduler) CheckIn(customer *customer.Customer) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if customer.IsVIP {
		waitTime := ServiceTime*len(s.VipQueue) + (ServiceTime * (len(s.VipQueue) / vipProcessingRate))
		s.VipQueue = append(s.VipQueue, customer)
		s.totalCheckedIn++
		customer.TicketNum = s.totalCheckedIn
		if waitTime == 0 {
			fmt.Printf("Thank you %s %s for checking in. You ticket number is %d. You are next in line to be seen.\n", customer.FirstName, customer.LastName, customer.TicketNum)
		} else {
			fmt.Printf("Thank you for %s %s for checking in. Your ticket number is %d. You have an approximate wait time of %d minutes\n",
				customer.FirstName, customer.LastName, customer.TicketNum, waitTime)
		}

	} else {
		waitTime := ServiceTime*len(s.RegularQueue) + (ServiceTime * (len(s.RegularQueue) * 2))
		s.RegularQueue = append(s.RegularQueue, customer)
		s.totalCheckedIn++
		customer.TicketNum = s.totalCheckedIn
		if waitTime == 0 {
			waitTime = 10
		}
		fmt.Printf("Thank you for %s %s for checking in. Your ticket number is: %d. You have an approximate wait time of %d minutes\n",
			customer.FirstName, customer.LastName, customer.TicketNum, waitTime)
	}
}

// Part 2: Retrieves the next customer to be served where all VIP customers are served first
// func (s *ServiceScheduler) GetNextCustomer() *customer.Customer {
// 	s.lock.Lock()
// 	defer s.lock.Unlock()
// 	var customer *customer.Customer

// 	if len(s.VipQueue) > 0 {
// 		customer = s.VipQueue[0]
// 		s.VipQueue = s.VipQueue[1:]
// 	} else {
// 		customer = s.RegularQueue[0]
// 		s.RegularQueue = s.RegularQueue[1:]
// 	}
// 	return customer
// }

// Part 3: Retrieves the next customer to be served while maintaining a 2:1 VIP to regular customer rate
func (s *ServiceScheduler) GetNextCustomer() *customer.Customer {
	s.lock.Lock()
	defer s.lock.Unlock()
	var customer *customer.Customer

	if len(s.VipQueue) > 0 && s.vipRate > 0 {
		customer = s.VipQueue[0]
		s.VipQueue = s.VipQueue[1:]
		s.vipRate--
	} else if len(s.RegularQueue) > 0 {
		customer = s.RegularQueue[0]
		s.RegularQueue = s.RegularQueue[1:]
		s.vipRate = vipProcessingRate
	} else if len(s.VipQueue) > 0 {
		customer = s.VipQueue[0]
		s.VipQueue = s.VipQueue[1:]
	}
	return customer
}
