package main

import (
	"fmt"

	id "../../id"
)

func main() {
	fmt.Println("multiverse sortable time stamped id")
	fmt.Println("============================================================")
	for i := 1; i <= 100; i++ {
		fmt.Println("Id.New().String():", id.New().String())
		fmt.Println("Id.New().Bytes():", id.New().Bytes())
		fmt.Println("Id.New().Nonce():", id.New().Nonce())
		fmt.Println("Id.New().Pid():", id.New().Pid())
		fmt.Println("Id.New().Time():", id.New().Time())
		fmt.Println("Id.New().Short():", id.New().Short())
		fmt.Println("Id.New().NoPrefix():", id.New().NoPrefix())
		v, _ := id.New().Value()
		fmt.Println("Id.New().Value():", v)
	}
}
