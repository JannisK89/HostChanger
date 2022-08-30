package main

import (
	"fmt"
	"os"
)

func chooseEnv(value int) []byte {

	if value == 1 {
		body, err := os.ReadFile("c:/Windows/System32/Drivers/etc/hosts_t1")
		if err == nil {
			return body
		}
	}
	if value == 2 {
		body, err := os.ReadFile("c:/Windows/System32/Drivers/etc/hosts_t2")
		if err == nil {
			return body
		}
	}
	if value == 3 {
		body, err := os.ReadFile("c:/Windows/System32/Drivers/etc/hosts_default")
		if err == nil {
			return body
		}
	}
	return []byte("No file found")

}

func main() {
	var userInput int
	fmt.Printf("Choose target environment:\n1. T1\n2. T2\n3. Prod\n")
	fmt.Scan(&userInput)

	if userInput < 4 && userInput > 0 {
		body := chooseEnv(userInput)
		err := os.WriteFile("c:/Windows/System32/Drivers/etc/hosts", body, 0744)
		println("Hostfile updated")

		if err != nil {
			fmt.Println("Error when writing to file")
			fmt.Println(err)
		}
	} else {
		println("Option not available")
	}
}
