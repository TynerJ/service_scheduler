package main

import (
	"bufio"
	"fmt"
	"os"
	"service_scheduler/customer"
	"service_scheduler/scheduler"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "service-scheduler",
	Short: "The service scheduler",
	Long:  `A CLI application for adding clients to the service scheduler.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the service scheduler",
	Run: func(cmd *cobra.Command, args []string) {
		s := scheduler.NewScheduler()
		reader := bufio.NewReader(os.Stdin)

		go func() {
			for {
				if len(s.RegularQueue) != 0 || len(s.VipQueue) != 0 {
					nextCustomer := s.GetNextCustomer()
					if nextCustomer != nil {
						fmt.Printf("Now serving customer: %d\n", nextCustomer.TicketNum)
						// The amount of time it would take to complete serving a customer
						// time.Sleep(time.Minute * serviceTime)
					} else {
						fmt.Println("No customers to serve.")
					}
					time.Sleep(time.Second * 30)
				}
			}
		}()
		for {
			fmt.Println("Do you wish to check in? (y/n)")
			input, _ := reader.ReadString('\n')
			input = strings.ToLower(input)[:len(input)-1]
			if input == "n" || input == "no" {
				os.Exit(0)
			} else if input != "y" && input != "yes" {
				fmt.Println("Invalid response. Please enter y or n")
			} else {
				fmt.Println("Please follow the following prompts: ")
				fmt.Println("Please enter your first name.")
				firstName, _ := reader.ReadString('\n')
				firstName = strings.TrimSpace(firstName)

				fmt.Println("Please enter your last name.")
				lastName, _ := reader.ReadString('\n')
				lastName = strings.TrimSpace(lastName)

				fmt.Println("Please enter your 10 digit phone number.")
				phoneNumber, _ := reader.ReadString('\n')
				phoneNumber = strings.TrimSpace(phoneNumber)

				customer, err := customer.NewCustomer(firstName, lastName, phoneNumber)
				if err != nil {
					fmt.Printf("%v\n", err)
				} else {
					s.CheckIn(customer)
					time.Sleep(time.Second * 2)
				}
			}
		}
	},
}

// If one wishes to use the service_scheduler without the CLI comment out all prior code above and uncomment all code below
// func main() {
// 	scheduler := scheduler.NewScheduler()

// 	// Customer check-ins
// 	customer1 := &customer.Customer{FirstName: "John", LastName: "Smith", PhoneNumber: "123-456-7890", IsVIP: false}
// 	customer2 := &customer.Customer{FirstName: "Alice", LastName: "Johnson", PhoneNumber: "987-654-3210", IsVIP: false}
// 	customer3 := &customer.Customer{FirstName: "Bob", LastName: "Meyers", PhoneNumber: "456-987-9638", IsVIP: false}
// 	customer4 := &customer.Customer{FirstName: "Mike", LastName: "Tomlin", PhoneNumber: "567-890-8527", IsVIP: true}
// 	customer5 := &customer.Customer{FirstName: "Clyde", LastName: "Jones", PhoneNumber: "245-753-0025", IsVIP: true}
// 	customer6 := &customer.Customer{FirstName: "Alex", LastName: "Jennings", PhoneNumber: "852-356-3468", IsVIP: false}
// 	customer7 := &customer.Customer{FirstName: "Jason", LastName: "Wright", PhoneNumber: "246-846-5132", IsVIP: true}
// 	customer8 := &customer.Customer{FirstName: "Carla", LastName: "Adams", PhoneNumber: "410-212-9987", IsVIP: true}
// 	customer9 := &customer.Customer{FirstName: "Gwen", LastName: "Tyler", PhoneNumber: "212-458-5867", IsVIP: false}
// 	customer0 := &customer.Customer{FirstName: "Virginia", LastName: "Jefferon", PhoneNumber: "918-361-1124", IsVIP: true}

// 	// Check-in customers
// 	scheduler.CheckIn(customer1)
// 	scheduler.CheckIn(customer2)
// 	scheduler.CheckIn(customer3)
// 	scheduler.CheckIn(customer4)
// 	scheduler.CheckIn(customer5)
// 	scheduler.CheckIn(customer6)
// 	scheduler.CheckIn(customer7)
// 	scheduler.CheckIn(customer8)
// 	scheduler.CheckIn(customer9)
// 	scheduler.CheckIn(customer0)

// 	// Process customers
// 	for i := 0; i < 9; i++ {
// 		nextCustomer := scheduler.GetNextCustomer()
// 		if nextCustomer != nil {
// 			fmt.Printf("Serving customer: %s %s\n", nextCustomer.FirstName, nextCustomer.LastName)
// 		} else {
// 			fmt.Println("No customers to serve.")
// 		}
// 	}
// }
