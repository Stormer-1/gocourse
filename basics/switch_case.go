package basics

import "fmt"

// CONCISE WAY TO WRITE MULTIPLE IF ELSE STATEMENTS. IT SIMPLIFIES THE CODE AND MAKES IT MORE READABLE.

// S E C V (Switch, expression, case, value) || S E C C (Switch, expression, case, condition)

func main() {

	// switch statement in go (switch case default) (fallthrough: executes the next case statement regardless of the condition)

	// switch expression {
	// case value1:
	// code to execute if condition (value1) equals expression or is true
	// fallthrough
	// case value2:
	// code to execute if condition (value2) equals expression or is true
	// case value3:
	// code to execute if condition (value3) equals expression or is true
	// default:
	// code to execute if expression does not match any value or none of the conditions are true
	// }

	// switch statement in other languages (switch case break default)

	// switch expression {
	// case value1:
	// code to execute if condition (value1) equals expression or is true
	// break
	// case value2:
	// code to execute if condition (value2) equals expression or is true
	// break
	// case value3:
	// code to execute if condition (value3) equals expression or is true
	// break
	// default:
	// code to execute if expression does not match any value or none of the conditions are true
	// }

// 	fruit := "pineapple"
// 
// 	switch fruit {
// 	case "apple":
// 		fmt.Println("apple")
// 	case "banana":
// 		fmt.Println("banana")
// 	default:
// 		fmt.Println("Unknown fruit!")
// 	}
// 
	// Multiple Conditions (multiple cases in a single case statement)
// 	day := "Monday"
// 
// 	switch day {
// 	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
// 		fmt.Println("Weekday")
// 	case "Saturday", "Sunday":
// 		fmt.Println("Weekend")
// 	default:
// 		fmt.Println("Another day probably, in the multiverse!")
// 	}
// 
	// Expressions (switch with no expression)
// 
// 	number := 15
// 
// 	switch {
// 	case number < 10:
// 		fmt.Println("Number is less than 10")
// 	case number >= 10 && number < 20:
// 		fmt.Println(number, "is between 10 and 19")
// 	default:
// 		fmt.Println(number, "is 20 or more")
// 	}
// 
// 	num:= 5
// 
// 	switch {
// 	case num > 1:
// 		fmt.Println("Number is greater than 1")
// 		fallthrough
// 	case num == 5:
// 		fmt.Println("Number is 5")
// 	default:
// 		fmt.Println("Number is less than 1")
// 	}

	checkType(10)
	checkType(10)
	checkType("Peace")
	checkType(3.14)
	checkType("Hello")
	checkType(true)
}

// interface{} is a type that can hold any value; it is used to store values of different types in a single variable
// No fallthrough in type switch 
// Type switch is used to check the type of a variable
func checkType(guess interface{}) {
	switch guess.(type) {
		case int:
			fmt.Println("guess is an integer")
		case int8:
			fmt.Println("guess is an int8")
		case float64:
			fmt.Println("guess is a float")
		case string:
			fmt.Println("guess is a string")
		default:
			fmt.Println("guess is of an unknown type")
	}
}