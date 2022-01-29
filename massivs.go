package main

import(
	"fmt"
)

func main() {
	var todoList = [3]string{
		"купить молоко",
		"купить хлеб",
		"купить печенье",
	}
	fmt.Printf("К-во элементов в списке %d\n", len(todoList))

	for index, item := range todoList {
		fmt.Printf("%d. %s\n", index, item)
	}

}