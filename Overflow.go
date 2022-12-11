package main

import "fmt"

// When we are trying to assign values greater than supported range, there is a possibility of the overflowing
// For Signed int, if overflow is happening - 2's complement is calculated and sign is changed
// For unsigned int, only the values present in the allocated memory bits is considered, overflow bits are ignored

func Overflow() {
	fmt.Println("Overflow function is called")
	var Vint8 int8 = 127 // -128 to 127 Signed integer
	// While initializing , program can not assign value more than the allowed range
	// ERROR:
	// var Vuint16 uint16 = 2345566777
	// cannot use 2345566777 (untyped int constant) as uint16 value in variable declaration (overflows)

	var Vuint16 uint16 = 65535 // For unsgined int, range is 0 to 65535
	var Vint64 int64 = 90
	Vuint16 = Vuint16 + 22

	// ERROR:
	// Vuint16 = Vuint16+100000 (untyped int constant 100000) overflows uint16
	// Even if constant value is being used in equation to assign the result to variable,
	// the value is checked if it is in range

	// ERROR:
	// Vuint16 += Vuint8 // Type mismatch

	fmt.Println(" Vint8 -", Vint8)
	fmt.Println(Vint8, Vint8+1, Vint8+2, Vint8+3, Vint8-127, Vint8-127-1, Vint8-127-2, Vint8+127, Vint8+127+2)

	var Vuint8 uint8 = 255
	fmt.Println(Vuint8, Vuint8+1, Vuint8+2, Vuint8+3, Vuint8-128, Vuint8-129, Vuint8-130, Vuint8-255, Vuint8-255-2)

	fmt.Println(" Vint64 -", Vint64)
	fmt.Println(Vint64, Vint64+1, Vint64+2, Vint64+3, Vint64-128, Vint64-129, Vint64-130, Vint64-255, Vint64-255-2)

	var Vbool bool = true
	fmt.Println(Vbool)
}
