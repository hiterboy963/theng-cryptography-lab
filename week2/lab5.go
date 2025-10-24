package main

import (
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func tobinary(b []byte) string {
	var sb strings.Builder
	for i := 0; i < len(b); i++ {
		sb.WriteString(fmt.Sprintf("%08b", b[i]))
	}
	return sb.String()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	input = strings.TrimSpace(input)

	data := []byte(input)

	fmt.Println("Original String:", input)
	fmt.Println("Hexadecimal:", hex.EncodeToString(data))
	fmt.Println("Base64:", base64.StdEncoding.EncodeToString(data))
	fmt.Println("Binary:", tobinary(data))
}
