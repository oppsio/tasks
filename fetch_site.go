package main

import "fmt"

type fetchSite struct{}

func (fetchSite) Run() {
	fmt.Println("fetch site is running!!")
}
