package main

import "fmt"

func main() {
	// const is a constant declaration
	// const is must initalize with a value
	// if const not used, it will be zero value
	const a = 1

	const firstName string = "Ario"

	const (
		lastName     = "Pratomo"
		age      int = 24
	)

	fmt.Println("Nama Depan = ", firstName)
	fmt.Println("Nama belakang = ", lastName)
	fmt.Println("Umur = ", age)

}
