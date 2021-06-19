// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Package Variable
//
//  1. Declare a variable in the package-scope
//
//  2. Observe whether something happens when you don't
//     use it
// ---------------------------------------------------------

var a int = 12
var b int = a + 3
func main() {
	fmt.Println("something")
	fmt.Println(a)
	fmt.Println(b)
}
