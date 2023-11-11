package main

import "Pz5/app"

func main() {
	err := app.Run()

	if err != nil {
		panic(err)
	}
}
