package main

import "fmt"

func main() {
	// variable is a storage for data
	// variable name is case sensitive
	// variable name must start with a letter or underscore
	// variable name can contain letters, numbers, and underscores
	// in golang, variable is for storing same type of data, and if you want to store different type of data, you need to declare another variable
	// for create variable, you can use var keyword

	var name string

	name = "Rio"
	fmt.Println("Nama = ", name)

	name = "Ario Pratomo"
	fmt.Println("Nama = ", name)

	var age = 24
	fmt.Println("Umur = ", age)

	country := "Indonesia"
	fmt.Println("Negara = ", country)

	// multiple variable declaration
	var (
		name2 = "Rio"
		age2  = 24
	)
	fmt.Println("Nama = ", name2)
	fmt.Println("Umur = ", age2)

}
