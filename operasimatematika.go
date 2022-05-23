package main

import "fmt"

func main() {

	// operator
	// + - * / %
	// ++ --
	// == != > < >= <=
	// && ||
	// & | ^ << >>
	// !

	var result = 10 + 5
	fmt.Println("10 + 5 = ", result)

	var a = 10
	var b = 5
	var c = a * b
	fmt.Println("10 * 5 = ", c)

	a += 10
	fmt.Println("a += 10 = ", a)
	a -= 10
	fmt.Println("a -= 10 = ", a)
	a *= 10
	fmt.Println("a *= 10 = ", a)
	a /= 10
	fmt.Println("a /= 10 = ", a)
	a %= 10
	fmt.Println("a %= 10 = ", a)

	// unary operator
	a++
	fmt.Println("a++ = ", a)
	a--
	fmt.Println("a-- = ", a)
	var positive = +a
	fmt.Println("+a = ", positive)
	var negative = -a
	fmt.Println("-a = ", negative)

	// operator ternary
	var x = 10
	var y = 5
	var z = x > y
	fmt.Println("z = ", z)

}
