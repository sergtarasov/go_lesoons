package main

import "fmt"

const pi = 3.1415



func main() {
	printCircleArea(2)
	printCircleArea(4)

	//calculateCircleArea := 5

	fmt.Println("Площадь круга с радиусом 5см=", calculateCircleArea(5))
}

func printCircleArea(radius int) {
	
	fmt.Printf("Радиус круга: %d см.\n", radius)
	fmt.Println("Формула для расчета площади круга: A=pir2\n")

	circleArea := calculateCircleArea(radius)
	fmt.Printf("Площадь круга: %f32 см кв.\n", circleArea)

}

func calculateCircleArea(radius int) float32 {

	return float32(radius) * float32(radius) * pi
}
