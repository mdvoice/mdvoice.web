package main

import "fmt"

func checknil(e error) {
	if e != nil {
		fmt.Printf("error! at %s\n", e.Error())
		panic(e)
	}
}
func checkstr(e string) {
	if e == "" {
		fmt.Printf("error! at '%s' is empty\n", e)
		panic(e)
	}
}
