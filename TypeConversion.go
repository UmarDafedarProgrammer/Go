/*
	Type Casting / Type Conversion
	--------------------------------------------------------
	1. Converting data from one type to another type is called type conversion
	2. Go is strong type System
	3. Go does not support implicit type conversion /automatic type conversion
		even if the data types are compatible
	4. Explicit type conversion is allowed between compatible data types
	5. Var_New_Data_type = New_data_type(Value/Variable)
	6. Assignment between two mixed types is not allowed
	7. Expressions with mixed types is not allowed like addition, subtraction etc.
*/
package main

import "fmt"

func TypeConversion() {
	fmt.Println("TypeConversion function is called")
	var v_float32 float32 = 3.1345

	// var v_int int = v_float32
	//cannot use v_float32 (variable of type float32) as type int in variable declaration

	var v_int int = int(v_float32)
	v_int32 := int(v_float32) // float to int
	fmt.Println(v_int)
	fmt.Println(v_int32)

	v_int = 3456
	v_float32 = float32(v_int) // int to float
	fmt.Println(-v_float32)

	var v_string string = "Hello"
	fmt.Println(v_string)

	v_string = string("�")
	fmt.Println("After type conversion ", v_string)
	v_int = int('�')
	fmt.Println("After type conversion ", v_int)
	//v_string = string(18989997) //Not compatible. Maximum ASCII value is 65533
	v_string = string(1101)
	fmt.Println("After type conversion ", v_string)
	v_int = int('�')
	fmt.Println("After type conversion ", v_int)
	fmt.Println("Binary Numbers")
	fmt.Printf("%b \n", 18989997)
	fmt.Printf("%b \n", v_int)
}
