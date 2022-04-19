package cars
import (
	"fmt"
	"strconv"
)

type ICar interface{
	create(index int)
}

type Volvo struct {}
func (v Volvo) create(index int) {
	formattedIndex := strconv.Itoa(index)
	fmt.Println(formattedIndex + ". Вольво собрана")
}

type Kia struct {}
func (k Kia) create(index int) {
	formattedIndex := strconv.Itoa(index)
	fmt.Println(formattedIndex + ". Kia собрана")
}

type Mercedes struct {}
func (m Mercedes) create(index int) {
	formattedIndex := strconv.Itoa(index)
	fmt.Println(formattedIndex + ". Mercedes собран")
}

func createCar(index int, car ICar) {
	car.create(index)
}

func CreateCars() {
	volvo := Volvo{}
	kia := Kia{}
	mercedeses := Mercedes{}
	cars := []ICar{volvo, kia, mercedeses}
	const count = 10

	for _, car := range cars {
		plannedCars := make([]ICar, count)
		
		for i, _ :=  range plannedCars{
			createCar(i, car)
		}
	}
}