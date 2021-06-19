// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Leap Year
//
//  Find out whether a given year is a leap year or not.
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me a year number
//
//  go run main.go eighties
//    "eighties" is not a valid year.
//
//  go run main.go 2018
//    2018 is not a leap year.
//
//  go run main.go 2100
//    2100 is not a leap year.
//
//  go run main.go 2019
//    2019 is not a leap year.
//
//  go run main.go 2020
//    2020 is a leap year.
//
//  go run main.go 2024
//    2024 is a leap year.
// ---------------------------------------------------------

func main() {

	if len(os.Args) <= 1 {
		fmt.Println("Give me a year number")
	}else if number,err := strconv.Atoi(os.Args[1]); err != nil || number == 0 {
		_ = number
		fmt.Println(number)
		fmt.Println(err)
		fmt.Printf("\"%s\" is not a valid year",os.Args[1])
	}else if float64(number)/100 == math.Ceil(float64(number)/100) && number % 400 == 0 {
		fmt.Println(number)
		fmt.Printf("%s is a leap year", os.Args[1])
	}else if float64(number)/4 == math.Ceil(float64(number)/4) && float64(number)/100 != math.Ceil(float64(number)/100) {
		fmt.Printf("%s is a leap year", os.Args[1])
	}else {
		fmt.Printf("%s is not a leap year",os.Args[1])
	}
}
