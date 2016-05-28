package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://api.phteven.io/"

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		return
	}

	// Defer is used to ensure that a function call is performed later in a program’s execution, usually for purposes of cleanup. defer is often used where e.g. ensure and finally would be used in other languages.
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading body: %v\n", err)
		return
	}

	fmt.Printf("Response: %s\n", body)
}
