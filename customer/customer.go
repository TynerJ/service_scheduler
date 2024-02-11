package customer

import (
	"errors"
	"math/rand"
	"regexp"
	"strings"
)

type Customer struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	IsVIP       bool
	TicketNum   int
}

// Validates the given parameters and creates a new customer if they are all valid
func NewCustomer(firstName, lastName, number string) (*Customer, error) {
	firstName = strings.ReplaceAll(firstName, " ", "")
	lastName = strings.ReplaceAll(lastName, " ", "")
	if ok := ValidatePhoneNumber(number); ok && firstName != "" && lastName != "" {
		return &Customer{
			FirstName:   firstName,
			LastName:    lastName,
			PhoneNumber: number,
			IsVIP:       VIPCheck(),
		}, nil
	} else {
		return nil, errors.New("Invalid first name, last name or phone number")
	}
}

func ValidatePhoneNumber(number string) bool {
	//Allows for the first three digits to be enclosed in parenthesis
	regex := `^\(?\d{3}\)?[-.\s]?\d{3}[-.\s]?\d{4}$`

	re := regexp.MustCompile(regex)

	// Check if the input string matches the regular expression
	return re.MatchString(number)
}

// Function to act as a database to check if a customer is a VIP or not.
func VIPCheck() bool {
	// Generate a random integer (0 or 1)
	randomNumber := rand.Intn(2)

	// Convert the random number to a boolean
	return randomNumber == 1
}
