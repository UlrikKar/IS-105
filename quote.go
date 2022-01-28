package main

import (
	"fmt"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Dette viser alle funksjonene")
	fmt.Println(quote.Opt())
	fmt.Println(quote.Glass())
	fmt.Println(quote.Go())
	fmt.Println(quote.Hello())
}
