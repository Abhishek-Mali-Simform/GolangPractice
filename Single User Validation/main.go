package main

import (
	"fmt"
	"os"
)

const (
	usage    = "Usage: [username] [password]"
	errUser  = "Access Denied for %q.\n"
	errPass  = " Incorrect Password for %q.\n"
	succUser = "Access Granted to %q.\n"

	usr  = "abhi"
	pass = "1999"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println(usage)
		return
	}
	u, p := args[1], args[2]
	if u != usr {
		fmt.Printf(errUser, u)
	} else if p != pass {
		fmt.Printf(errPass, u)
	} else {
		fmt.Printf(succUser, u)
	}
}
