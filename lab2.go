// Lab 2 code

package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Enter two integers(a & b): ")
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println("Invalid input:", err)
		return
	}
	fmt.Printf("You entered: a = %d, b = %d\n", a, b)

	bothpositive := (a > 0) && (b > 0)
	onegreater := (a > b) || (b > a)
	notequal := !(a == b)

	fmt.Printf("Both a and b are positive: %t\n", bothpositive)
	fmt.Printf("At least one of a or b is greater: %t\n", onegreater)
	fmt.Printf("a and b are not equal: %t\n", notequal)
}
