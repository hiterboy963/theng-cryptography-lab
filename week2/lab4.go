package main

import (
	"errors"
	"fmt"
)

func add(a, b int) int { return a + b }
func sub(a, b int) int { return a - b }
func mul(a, b int) int { return a * b }
func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero is null")
	}
	return a / b, nil
}
func mod(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("modulus by zero is null")
	}
	return a % b, nil
}

func main() {
	for {
		fmt.Println("MINI CALCULATOR")
		fmt.Println("1.Add 2.Sub 3.Mul 4.Div 5.Mod 6.Exit")
		fmt.Print("choose:")

		var choice int
		if _, err := fmt.Scan(&choice); err != nil {
			fmt.Println("Invalid input:", err)
			fmt.Scanln()
			continue
		}

		if choice == 6 {
			fmt.Println("Exiting...")
			break
		}
		if choice < 1 || choice > 6 {
			fmt.Println("Invalid choice, please select a valid option.")
			continue
		}

		var a, b int
		fmt.Print("Enter two integers(a & b): ")
		if _, err := fmt.Scan(&a, &b); err != nil {
			fmt.Println("Invalid input:", err)
			continue
		}

		switch choice {
		case 1:
			fmt.Printf("Result: %d\n", add(a, b))
		case 2:
			fmt.Printf("Result: %d\n", sub(a, b))
		case 3:
			fmt.Printf("Result: %d\n", mul(a, b))
		case 4:
			if result, err := div(a, b); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Result: %d\n", result)
			}
		case 5:
			if result, err := mod(a, b); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Result: %d\n", result)
			}
		}
	}
}
