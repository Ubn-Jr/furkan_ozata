package main

import "fmt"

func main() {
	math(8)

}

func math(x int) {

	if x%2 == 0 {
		fmt.Println("sayi cift")
	} else {
		fmt.Println("sayi tek")
	}

}
