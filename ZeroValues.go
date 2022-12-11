package main

import "fmt"

/* Zero Value --> Default Value
By default for all the variables, which are not initialized. Go assigns a default value called zero value
For integers, zero value is 0
For strings, zero value is ""
For boo, it is false
*/

func ZeroValue() {
	fmt.Println("LOG : ZeroValue function is called")

	var v_int8 int8
	var v_int16 int16
	var v_int32 int32
	var v_int64 int64
	var v_int int

	fmt.Printf(`
	v_int8		:%d	%T
	v_int16		:%d	%T
	v_int32		:%d	%T
	v_int64		:%d	%T
	v_int		:%d	%T
	`, v_int8, v_int8, v_int16, v_int16, v_int32, v_int32, v_int64, v_int64, v_int, v_int)

	fmt.Println()
	var v_uint8 uint8
	var v_uint16 uint16
	var v_uint32 uint32
	var v_uint64 uint64
	var v_uint uint

	fmt.Printf(`
	v_uint8		:%d	%T
	v_uint16	:%d	%T
	v_uint32	:%d	%T
	v_uint64	:%d	%T
	v_uint		:%d	%T
	`, v_uint8, v_uint8, v_uint16, v_uint16, v_uint32, v_uint32, v_uint64, v_uint64, v_uint, v_uint)

	fmt.Println()
	var v_bool bool
	fmt.Printf(`
	v_bool		:%t	%T`, v_bool, v_bool)

	fmt.Println()
	var v_string string
	fmt.Printf(`
	v_string	:%s %T`, v_string, v_string)

	fmt.Println()
	var v_float64 float64
	var v_float32 float32
	fmt.Printf(`
	v_float64		:%.16f	%T
	v_float32		:%f	%T
	`, v_float64, v_float64, v_float32, v_float32)

	fmt.Println()
	var v_complex64 complex64
	var v_complex128 complex128
	fmt.Printf(`
	v_complex64		:%f	%T
	v_complex128		:%f	%T
	`, v_complex64, v_complex64, v_complex128, v_complex128)

	fmt.Println()
	var v_rune rune
	var v_uintptr uintptr
	fmt.Printf(`
	v_rune		:%d	%T
	v_uintptr		:%d	%T
	`, v_rune, v_rune, v_uintptr, v_uintptr)
}
