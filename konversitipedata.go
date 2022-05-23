package main

import "fmt"

func main() {

	// convert to boolean
	var i32 int32 = 32768
	var i64 int64 = int64(i32)
	var i8 int8 = int8(i32)

	fmt.Println("i32 = ", i32)
	fmt.Println("i64 = ", i64)
	fmt.Println("i8 = ", i8)
	fmt.Println("==========================")
	var name = "Ario Pratomo"
	var a = name[0]
	var aString = string(a)
	fmt.Println("name = ", name)
	fmt.Println("a = ", a)
	fmt.Println("aString = ", aString)

}
