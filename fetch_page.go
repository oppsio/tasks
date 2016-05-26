package main

import "fmt"

type fetchPage struct{}

func (fetchPage) Run() {
	fmt.Println("fetch.page is running!!")
}
