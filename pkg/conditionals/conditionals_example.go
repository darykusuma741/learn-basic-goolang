package conditionals

import "fmt"

func ConditionalsExample() {
	// ==============================
	// IF with short statement
	// ==============================
	if example := 15; example%2 == 0 {
		fmt.Println("Example is even (short statement in if)")
	} else {
		fmt.Println("Example is odd (short statement in if)")
	}

	// ==============================
	// Nested IF
	// ==============================
	score := 85
	if score >= 60 {
		if score >= 90 {
			fmt.Println("Excellent!")
		} else if score >= 75 {
			fmt.Println("Good!")
		} else {
			fmt.Println("Passed")
		}
	} else {
		fmt.Println("Failed")
	}

	// ==============================
	// SWITCH with multiple cases
	// ==============================
	fruit := "apple"
	switch fruit {
	case "apple", "pear", "peach":
		fmt.Println("This is a common fruit")
	case "dragonfruit", "durian":
		fmt.Println("This is an exotic fruit")
	default:
		fmt.Println("Unknown fruit")
	}

	// ==============================
	// SWITCH without expression (like if-else)
	// ==============================
	number := 42
	switch {
	case number < 0:
		fmt.Println("Negative")
	case number%2 == 0:
		fmt.Println("Even number")
	default:
		fmt.Println("Odd number")
	}

	// ==============================
	// Type switch
	// ==============================
	var example interface{} = 3.14

	switch v := example.(type) {
	case int:
		fmt.Println("Integer:", v)
	case float64:
		fmt.Println("Float:", v)
	case string:
		fmt.Println("String:", v)
	default:
		fmt.Println("Unknown type")
	}
}
