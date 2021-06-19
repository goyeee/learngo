// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

func main() {
	const min int = 1
	const pi float64 = 3.14
	const version string = "2.0.1"
	const debug bool = true
	const A int8 = 'A'

	fmt.Println(min, pi, version, debug, A)

	var a rune
	var b int32
	a = 32
	b = a + 1
	fmt.Println(b)
}
