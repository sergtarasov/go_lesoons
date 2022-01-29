package main

import (
	"fmt"
)

func main() {
	x := 10
	
	fmt.Println("значение x:", x)

	increment(&x)

	fmt.Println("значение x после вызова функции increment():", x)
}

func increment(p *int){
	*p += 1
}