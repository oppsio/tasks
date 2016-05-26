package main

import "fmt"

type parsePage struct{}

func (parsePage) Run() {
	fmt.Println("parse.page is running!!")
}
