package basics

import "fmt"

type Employee struct {
	FirstName string
	LastName string
	Age int
}



func main() {

	const MAXRETRIES = 5 

	var employeeID = 1001
	
	fmt.Println("Employee ID: ", employeeID)

/*
PascalCase
use case for Structs, interfaces, enums, etc.
E.g. PersonName, CarName, HouseName

snake_case
use case for variables
E.g. first_name, last_name, full_name

UPPERCASE
use case for constants
E.g. PI, MAX_VALUE, MIN_VALUE, URL, API_KEY

camelCase
use case for variables
E.g. firstName, lastName, fullName


mixedCase
use case for variables
E.g. firstName, lastName, fullName

kebab-case
use case for variables
E.g. first-name, last-name, full-name

SCREAMING_SNAKE_CASE
use case for constants
*/
}