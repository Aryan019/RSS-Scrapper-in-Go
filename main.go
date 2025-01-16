package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello from the go server ")
	godotenv.Load()

	// Reading in the port from the file we have
	portString := os.Getenv("PORT")

	if portString == "" {
		fmt.Println("PORT not found in the environment variables")
		return
	}

	fmt.Println("Port -> ", portString)

}
