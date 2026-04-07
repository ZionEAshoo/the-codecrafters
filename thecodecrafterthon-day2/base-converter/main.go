package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func printHelp() {
	fmt.Println("Usage:")
	fmt.Println("  convert <value> <base>")
	fmt.Println("")
	fmt.Println("Bases:")
	fmt.Println("  bin → binary")
	fmt.Println("  dec → decimal")
	fmt.Println("  hex → hexadecimal")
	fmt.Println("")
	fmt.Println("Examples:")
	fmt.Println("  convert 1E hex")
	fmt.Println("  convert 10 bin")
	fmt.Println("  convert 255 dec")
	fmt.Println("")
	fmt.Println("Type 'quit' to exit.")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("🔢 Base Converter (type 'help' for usage)")

	for {
		fmt.Print("> ")

		if !scanner.Scan() {
			fmt.Println("\nGoodbye")
			break
		}

		input := strings.TrimSpace(scanner.Text())

		if input == "" {
			continue
		}

		fields := strings.Fields(input)

		command := fields[0]

		if command == "quit" {
			fmt.Println("Goodbye")
			break
		}

		if command == "help" {
			printHelp()
			continue
		}

		if command != "convert" {
			fmt.Println("Unknown command. Type 'help'.")
			continue
		}

		if len(fields) != 3 {
			fmt.Println("Error: Usage → convert <value> <base>")
			continue
		}

		value := fields[1]
		base := strings.ToLower(fields[2])

		switch base {

		case "dec":

			num, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				fmt.Println("Error: Invalid decimal number.")
				continue
			}

			bin := strconv.FormatInt(num, 2)
			hex := strings.ToUpper(strconv.FormatInt(num, 16))

			fmt.Printf("Binary:  %s\n", bin)
			fmt.Printf("Hex:     %s\n", hex)

		case "bin":

			num, err := strconv.ParseInt(value, 2, 64)
			if err != nil {
				fmt.Println("Error: Invalid binary number.")
				continue
			}

			fmt.Printf("Decimal: %d\n", num)

		case "hex":

			num, err := strconv.ParseInt(value, 16, 64)
			if err != nil {
				fmt.Println("Error: Invalid hexadecimal number.")
				continue
			}

			fmt.Printf("Decimal: %d\n", num)

		default:
			fmt.Println("Error: Base must be 'bin', 'dec', or 'hex'.")
		}
	}
}
