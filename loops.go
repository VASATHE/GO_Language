package main

import "fmt"

func loops() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
