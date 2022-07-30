package main

import "fmt"

func main() {
	onaIkiyeBolunebilme(24)

}

func onaIkiyeBolunebilme(x int) {

	if x%4 == 0 && x%3 == 0 {
		fmt.Println("sayi 12' ye tam bolunur")
	} else {
		fmt.Println("sayi 12' ye tam bolunmez")
	}

}
