package main

import (
	"fmt"
)

func main() {
	// Declare variable
	url := "http://api.phteven.io/"

	// Package fmt implements formatted I/O with functions analogous to C's printf and scanf. The format 'verbs' are derived from C's but are simpler.
	fmt.Printf("Our API url is %s\n", url)
}
