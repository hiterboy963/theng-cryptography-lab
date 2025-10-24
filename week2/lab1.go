// Lab 1 code

package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Enter two integers(a & b): ")
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	fmt.Printf("You entered: a = %d, b = %d\n", a, b)

	// Operator '='
	x, y := a, b
	fmt.Printf("result operation x = a, y = b: x = %d, y = %d\n", x, y)
	// Operator '+='
	a += b
	fmt.Printf("result operation a += b: a = %d\n", a)
	// operator '-='
	a -= b
	fmt.Printf("result operation a -= b: a = %d\n", a)
	// operator '*='
	a *= b
	fmt.Printf("result operation a *= b: a = %d\n", a)
	// operator '/='
	if b != 0 {
		a /= b
		fmt.Printf("result operation a /= b: a = %d\n", a)
	} else {
		fmt.Println("Division by zero is not allowed.")
	}
	// operator '%='
	if b != 0 {
		a %= b
		fmt.Printf("result operation a %%= b: a = %d\n", a)
	} else {
		fmt.Println("Modulo by zero is not allowed.")
	}
}
