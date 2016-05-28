package main

import (
	"fmt"
	"net/http"
)

func main() {
	url := "http://api.phteven.io/"

	// Go uses error values to indicate an abnormal state.
	// Making errors part of the code is clearer than throwing exception.
	_, err := http.Get("test")
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		return
	}

	fmt.Printf("Our API url is %s\n", url)
}
