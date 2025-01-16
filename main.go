package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello from the go server ")

	// Reading in the port from the file we have
	portString := os.Getenv("PORT")

	if portString == "" {
		fmt.Println("PORT not found in the environment variables")
		return
	}

	fmt.Println("Port -> ", portString)

}
