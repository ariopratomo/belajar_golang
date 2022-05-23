package main

import "fmt"

func main() {
	// int8 is a signed 8-bit integer min: -128 max: 127
	var i8 int8 = 127
	fmt.Printf("i8: %d\n", i8)
	// int16 is a signed 16-bit integer min: -32768 max: 32767
	var i16 int16 = 32767
	fmt.Printf("i16: %d\n", i16)
	// int32 is a signed 32-bit integer min: -2147483648 max: 2147483647
	var i32 int32 = 2147483647
	fmt.Printf("i32: %d\n", i32)
	// int64 is a signed 64-bit integer min: -9223372036854775808 max: 9223372036854775807
	var i64 int64 = 9223372036854775807
	fmt.Printf("i64: %d\n", i64)
	// uint8 is an unsigned 8-bit integer min: 0 max: 255
	var ui8 uint8 = 255
	fmt.Printf("ui8: %d\n", ui8)
	// uint16 is an unsigned 16-bit integer min: 0 max: 65535
	var ui16 uint16 = 65535
	fmt.Printf("ui16: %d\n", ui16)
	// uint32 is an unsigned 32-bit integer min: 0 max: 4294967295
	var ui32 uint32 = 4294967295
	fmt.Printf("ui32: %d\n", ui32)
	// uint64 is an unsigned 64-bit integer min: 0 max: 18446744073709551615
	var ui64 uint64 = 18446744073709551615
	fmt.Printf("ui64: %d\n", ui64)
	// float32 is a 32-bit floating point number min: -3.402823e+38 max: 3.402823e+38
	var f32 float32 = 3.402823e+38
	fmt.Printf("f32: %f\n", f32)
	// float64 is a 64-bit floating point number min: -1.7976931348623157e+308 max: 1.7976931348623157e+308
	var f64 float64 = 1.7976931348623157e+308
	fmt.Printf("f64: %f\n", f64)
	// complex64 is a complex number with float32 real and imaginary parts min: -3.402823e+38 max: 3.402823e+38
	var c64 complex64 = 3.402823e+38
	fmt.Printf("c64: %f\n", c64)
	// complex128 is a complex number with float64 real and imaginary parts min: -1.7976931348623157e+308 max: 1.7976931348623157e+308
	var c128 complex128 = 1.7976931348623157e+308
	fmt.Printf("c128: %f\n", c128)
	// byte is an alias for uint8
	var b byte = 255
	fmt.Printf("b: %d\n", b)
	// rune is an alias for int32
	var r rune = 65535
	fmt.Printf("r: %d\n", r)
	// int is minimal int32
	var i int = 2147483647
	fmt.Printf("i: %d\n", i)
	// uint is minimal uint32
	var u uint = 4294967295
	fmt.Printf("u: %d\n", u)

}
