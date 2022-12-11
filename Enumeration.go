package main

import "fmt"

/*
	Enum is used to define a data type that has set of named related values called elements, enumerals, enumerators
	It will have set of values, which are considered valid.
	in Go, there is no concept called enums.instead we use const to define enums

Go does not have enumerate types
Instead, you can use the special name iota in a single
const declaration to get a series of increasing values.

*/

func Enum() {
	fmt.Println("Enum function is called")

	/*	const (
				a
				b
				c
			)

		ERROR:
			Missing init*/

	// fmt.Println(iota)
	// ERROR: cannot use iota outside constant declaration

	const (
		a = iota
		b
		c
	)
	fmt.Println(a, b, c) // 0 1 2

	const (
		d = iota + 1
		e
		f
	)
	fmt.Println(d, e, f) // 1 2 3

	const (
		g = 1 + iota
		h
		i
	)
	fmt.Println(g, h, i) // 1 2 3

	const (
		j = 1 + iota
		k = 2 + iota
		l
	)
	fmt.Println(j, k, l) // 1 3 4

	const (
		m = 1
		n
		o
	)
	fmt.Println(m, n, o) //1 1 1

	const (
		p = iota
		q = iota
		r = iota
	)
	fmt.Println(p, q, r) // 0 1 2

	const (
		s = 0 - iota
		t
		u
	)
	fmt.Println(s, t, u) // 0 -1 -2

	const (
		v = 10 - iota
		_ // blank identifier to skip a value in a list of constants.
		w
		x
	)
	fmt.Println(v, w, x) // 10 8 7

	// Multiple iotas in a single line
	const (
		q1, q2, q3 = iota, iota, iota
		r1, r2, r3
		s1, s2, s3
	)

	fmt.Printf("q1: %d, q2: %d, q3: %d r1: %d, r2: %d, r3: %d s1: %d, s2: %d, s3: %d", q1, q2, q3, r1, r2, r3, s1, s2, s3)

	// multi-line strings - `` quotes
	fmt.Printf(`
	q1: %d, q2: %d, q3: %d 
	r1: %d, r2: %d, r3: %d 
	s1: %d, s2: %d, s3: %d`, q1, q2, q3, r1, r2, r3, s1, s2, s3)

	const (
		Sunday int = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Printf(`
	Sunday		:%d
	Monday		:%d
	Tuesday		:%d
	Wednesday	:%d
	Thursday	:%d
	Friday		:%d
	Saturday	:%d`, Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)

	const (
		January  int = 1 + iota //iota = 0
		February                //iota = 1
		March                   //iota = 2
		April    int = 10       //iota = 3
		May          = 1 - iota //iota = 4
		June                    //iota = 5
	)
	fmt.Println("January  :", January)  // 1
	fmt.Println("February :", February) // 2
	fmt.Println("March    :", March)    // 3
	fmt.Println("April    :", April)    // 10
	fmt.Println("May      :", May)      // -3
	fmt.Println("June     :", June)     // -4
}
