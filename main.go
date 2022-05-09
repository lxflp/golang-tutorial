package main

import (
	"fmt"
	"net/http"
)

var age int

func main() {
	g := func() {
		fmt.Println("function g")
	}
	func() {
		fmt.Println("anonymous func")
	}()
	age = 99
	func(i int) {
		fmt.Println("anonymous func", age)
	}(age)
	fmt.Println("test")
	g()
	f := func(i, y int) {
		fmt.Println("function ", i, y)
	}
	f(9, 0)

	http.HandleFunc("/", ttest)

	http.ListenAndServe(":8990", nil)
}

func ttest(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "test %d", age)
}
