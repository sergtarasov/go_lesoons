package main
import(
	"fmt"
	"errors"
)

const pi = 3.1415

func main(){
	printCircleArea(-2)
}

func printCircleArea(radius int){
	circleArea, errors := calculateCircleArea(radius)
	if errors != nil{
		fmt.Println(errors.Error())
		return
	}

	fmt.Printf("Radius circle: %d cm.\n", radius)
	fmt.Println("Formula rascheta ploshady kruga: A=pir2.\n ")

	fmt.Printf("Плошадь круга: %f32 cm. sq.\n\n", circleArea)
}

func calculateCircleArea(radius int) (float32, error) {
	if radius <=0 {
		return float32(0), errors.New("Радиус круга не может отрицательным")
	}

	return float32(radius) * float32(radius) * pi, nil
}