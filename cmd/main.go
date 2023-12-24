package main

import (
	"Pz5/app"
	"fmt"
	"log"
)

func main() {
	err := app.Run()

	if err != nil {
		panic(err)
	}

	fmt.Print("Натисніть Enter для виходу...")
	_, err = fmt.Scanln()
	if err != nil {
		log.Fatal(err)
	}

}
