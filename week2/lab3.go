package main

import "fmt"

func andop(a, b int) int           { return a & b }
func orop(a, b int) int            { return a | b }
func xorop(a, b int) int           { return a ^ b }
func notop(a int) int              { return ^a }
func leftshift(a int, n uint) int  { return a << n }
func rightshift(a int, n uint) int { return a >> n }

func main() {
	var a, b int
	var shiftcount int

	fmt.Print("Enter two integers(a & b) and shift count(non negative): ")
	_, err := fmt.Scan(&a, &b, &shiftcount)
	if err != nil {
		fmt.Println("Invalid input:", err)
		return
	}
	if shiftcount < 0 {
		fmt.Println("Shift count must be greater than or equal to 0")
		return
	}
	sh := uint(shiftcount)

	fmt.Printf("You entered: a = %d, b = %d, shift count = %d\n\n", a, b, shiftcount)
	fmt.Printf("a AND b = %d\n", andop(a, b))
	fmt.Printf("a OR b = %d\n", orop(a, b))
	fmt.Printf("a XOR b = %d\n", xorop(a, b))
	fmt.Printf("NOT a = %d\n", notop(a))
	fmt.Printf("NOT b = %d\n", notop(b))
	fmt.Printf("a left shifted by %d = %d\n", shiftcount, leftshift(a, sh))
	fmt.Printf("b left shifted by %d = %d\n", shiftcount, leftshift(b, sh))
	fmt.Printf("a right shifted by %d = %d\n", shiftcount, rightshift(a, sh))
	fmt.Printf("b right shifted by %d = %d\n", shiftcount, rightshift(b, sh))
}
