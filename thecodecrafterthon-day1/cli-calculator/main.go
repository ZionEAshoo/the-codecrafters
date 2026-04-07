package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func printHelp() {
	fmt.Println("Available commands:")
	fmt.Println("  add <a> <b>  → addition")
	fmt.Println("  sub <a> <b>  → subtraction")
	fmt.Println("  mul <a> <b>  → multiplication")
	fmt.Println("  div <a> <b>  → division")
	fmt.Println("  help         → show this help message")
	fmt.Println("  quit         → exit the calculator")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to CLI Calculator")
	fmt.Println("Type 'help' for options")

	for {
		fmt.Print("> ")

		if !scanner.Scan() {
			fmt.Println("Goodbye")
			break
		}

		input := scanner.Text()
		fields := strings.Fields(input)

		if len(fields) == 0 {
			continue
		}

		command := fields[0]

		if command == "quit" {
			fmt.Println("Goodbye")
			break
		}

		if command == "help" {
			printHelp()
			continue
		}

		if len(fields) != 3 {
			fmt.Println("Error: Expected exactly 2 arguments. Try 'help'.")
			continue
		}

		a, err1 := strconv.Atoi(fields[1])
		b, err2 := strconv.Atoi(fields[2])

		if err1 != nil || err2 != nil {
			fmt.Println("Error: Arguments must be valid integers.")
			continue
		}

		switch command {
		case "add":
			fmt.Printf("✦ Result: %d\n", a+b)

		case "sub":
			fmt.Printf("✦ Result: %d\n", a-b)

		case "mul":
			fmt.Printf("✦ Result: %d\n", a*b)

		case "div":
			if b == 0 {
				fmt.Println("Error: Division by zero is not allowed.")
				continue
			}
			fmt.Printf("✦ Result: %d\n", a/b)

		default:
			fmt.Println("Unknown command. Type 'help' for available commands.")
		}
	}
}
