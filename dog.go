package main

type Dog struct {
	name  string
	age   int
	motto string
}

func (dog Dog) say() string {
	return dog.motto
}
