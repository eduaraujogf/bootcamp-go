package main

import "fmt"

func main() {
	condition_1, condition_2 := true, false

	if condition_1 != true {
		// instructions to be executed if the condition is true
	}

	if condition_1 != true {
		// instructions to be executed if the condition is true
	} else {
		// instructions to be executed if the condition is false
	}

	if condition_1 {
		// instructions to be executed if the condition is true
	} else if condition_2 {
		// instructions to be executed if the condition is tru
	} else {
		// instructions to be executed if the condition is false
	}

	//short if

	if declaration := f(); declaration == true {
		// instruction to be executed if the condition is true
	}

	days := 1
	switch days {
	case 0:
		fmt.Println("Monday")
	case 1:
		fmt.Println("Tuesday")
	case 2:
		fmt.Println("Wednesday")
	case 3:
		fmt.Println("Thursday")
	case 4:
		fmt.Println("Friday")
	case 5:
		fmt.Println("Saturday")
	case 6:
		fmt.Println("Sunday")
	default:
		fmt.Println("Unknown")
	}

	var age uint8 = 18

	switch {
	case age >= 150:
		fmt.Println("Are you Immortal?")
	case age >= 18:
		fmt.Println("He/She is of age.")
	default:
		fmt.Println("He/She is under age.")
	}

	daySwitch := "sunday"

	switch daySwitch {
	case "monday", "tuesday", "wednesday", "thursday", "friday":
		fmt.Printf("%s is a weekday\n", daySwitch)
	default:
		fmt.Printf("%s is a weekend\n", daySwitch)
	}

	switch dayShort := "monday"; dayShort {
	case "monday", "tuesday", "wednesday", "thursday", "friday":
		fmt.Printf("%s is a weekday\n", dayShort)
	default:
		fmt.Printf("%s is a weekend\n", dayShort)
	}

}

func f() bool {
	return true
}
