package main

import (
	"fmt"
	"log"

	"github.com/pionier2027/smgt/core"
)

func main() {
	if err := core.SetupDB(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database created successfully.")
}
