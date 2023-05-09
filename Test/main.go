package main

import "fmt"

func main() {
	A := [3]int{1, 2, 3}
	B := A[1:]
	C := B[2:]
	for _, v := range B {
		fmt.Println(v)
	}
	for _, v := range C {
		fmt.Println(v)
	}
}
