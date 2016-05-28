package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	body, _ := getPhtevenResponseBody()
	fmt.Printf("Response: %s\n", body)
}

func getPhtevenResponseBody() ([]byte, error) {
	url := "http://api.phteven.io/translate/"
	dog := Dog{"Phteven", 3, "Stephen is a silly sausage who drinks at starbucks"}

	requestBody := bytes.NewBufferString(fmt.Sprintf("text=%s", dog.say()))
	response, err := http.Post(url, "application/x-www-form-urlencoded", requestBody)
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		return []byte(""), err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading body: %v\n", err)
		return []byte(""), err
	}

	return body, nil
}
