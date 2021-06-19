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
	"os"
	"strings"
)

// ---------------------------------------------------------
// CHALLENGE #1
//  Create a user/password protected program.
//
// EXAMPLE USER
//  username: jack
//  password: 1888
//
// EXPECTED OUTPUT
//  go run main.go
//    Usage: [username] [password]
//
//  go run main.go albert
//    Usage: [username] [password]
//
//  go run main.go hacker 42
//    Access denied for "hacker".
//
//  go run main.go jack 6475
//    Invalid password for "jack".
//
//  go run main.go jack 1888
//    Access granted to "jack".
// ---------------------------------------------------------
const usage = `
		Usage: [username] [password]
`
const (
	username = "jack"
	password = "1888"
	accessDenied = "Access denied for %q.\n"
	invalidPassword = "Invalid password for %q.\n"
	accessGranted = "Access granted to %q.\n"
)
func main() {


	if len(os.Args) < 3 {
		fmt.Println(strings.TrimSpace(usage))
		return
	}

	u, p := os.Args[1],os.Args[2]

	if u != username {
		fmt.Printf(accessDenied, u)
	}else if p != password {
		fmt.Printf(invalidPassword, u)
	}else {
		fmt.Printf(accessGranted,u)
	}
}
