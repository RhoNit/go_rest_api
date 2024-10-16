package main

import "fmt"

func Run() error {
	fmt.Println("Application is running...")
	return nil
}

func main() {
	fmt.Println("Server started...")
	if err := Run(); err != nil {
		fmt.Printf("Error while running the application: %s", err.Error())
	}
}
