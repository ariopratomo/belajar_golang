package main

import "fmt"

func main() {

	type NoKtp string
	type Married bool
	var noktpRio NoKtp = "123456789"
	var isMarried Married = false
	fmt.Println("No KTP Rio = ", noktpRio)
	fmt.Println("Is Married Rio = ", isMarried)

}
