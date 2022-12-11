package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func Const() {
	fmt.Println("Const Function is called")

	// while declaring const var should not be used
	// const var Cint int = 30

	// walrus operator should not be used for declaring constant
	// const CString := "Hello"

	// Typed Constant
	const Cvar int = 210

	// Once initialized, the value can not be changed for constant
	// Cvar = 30

	fmt.Println("CVar value :", Cvar)
	fmt.Println("Type Of CVar value :", reflect.TypeOf(Cvar))
	fmt.Println("Type of CVar :", reflect.TypeOf(Cvar).String())
	fmt.Println("Kind of CVar :", reflect.TypeOf(Cvar).Kind())
	fmt.Println("Value of CVar :", reflect.ValueOf(Cvar))

	// untyped constant
	const n = 3000000000000
	fmt.Println("n value :", n)
	fmt.Println("Type Of n value :", reflect.TypeOf(n))
	fmt.Println("Type of n :", reflect.TypeOf(n).String())
	fmt.Println("Kind of n :", reflect.TypeOf(n).Kind())
	fmt.Println("Value of n :", reflect.ValueOf(n))
	fmt.Println("Size of n :", unsafe.Sizeof(n))

	const m = n / 3.0
	fmt.Println("m value :", m)
	fmt.Println("Type Of m value :", reflect.TypeOf(m))
	fmt.Println("Type of m :", reflect.TypeOf(m).String())
	fmt.Println("Kind of m :", reflect.TypeOf(m).Kind())
	fmt.Println("Value of m :", reflect.ValueOf(m))
	fmt.Println("Size of m :", unsafe.Sizeof(m))

	const str string = "UmarDafedar"
	fmt.Println("str value :", str)
	fmt.Println("Type Of str value :", reflect.TypeOf(str))
	fmt.Println("Type of str :", reflect.TypeOf(str).String())
	fmt.Println("Kind of str :", reflect.TypeOf(str).Kind())
	fmt.Println("Value of str :", reflect.ValueOf(str))
	fmt.Println("Size of str :", unsafe.Sizeof(str))

	const Ready bool = true
	fmt.Println("Are you ready ?:", Ready)
	fmt.Println("Type Of Ready value :", reflect.TypeOf(Ready))
	fmt.Println("Type of Ready :", reflect.TypeOf(Ready).String())
	fmt.Println("Kind of Ready :", reflect.TypeOf(Ready).Kind())
	fmt.Println("Value of Ready :", reflect.ValueOf(Ready))
	fmt.Println("Size of Ready :", unsafe.Sizeof(Ready))

	const UnnamedConstant = m / n
	fmt.Println("UnnamedConstant:", UnnamedConstant)
	fmt.Println("Type Of UnnamedConstant value :", reflect.TypeOf(UnnamedConstant))
	fmt.Println("Type of UnnamedConstant :", reflect.TypeOf(UnnamedConstant).String())
	fmt.Println("Kind of UnnamedConstant :", reflect.TypeOf(UnnamedConstant).Kind())
	fmt.Println("Value of UnnamedConstant :", reflect.ValueOf(UnnamedConstant))
	fmt.Println("Size of UnnamedConstant :", unsafe.Sizeof(UnnamedConstant))

	// An untyped constant has no limits. Depends up on the value assigned type will be decided

	var reminder = 1234 / UnnamedConstant
	fmt.Println("reminder:", reminder)
	fmt.Println("Type Of reminder value :", reflect.TypeOf(reminder))
}
