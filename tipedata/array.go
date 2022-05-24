package main

import "fmt"

func main() {
	var name [2]string
	name[0] = "Ario"
	name[1] = "Pratomo"
	fmt.Println(name[0], name[1])

	var values = [5]int{1, 2, 3, 4, 5}
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
	fmt.Println(values[3])
	fmt.Println(values[4])

	length := len(values)
	fmt.Println(length)
}
