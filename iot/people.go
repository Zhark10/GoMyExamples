package iot

import (
	"fmt"
	"os"
)

type person struct {
	name   string
	age    int32
	weight float64
}

func IotPeople() {
	people := [3]person{
		{
			name:   "Arkady",
			age:    27,
			weight: 63,
		},
		{
			name:   "Dasha",
			age:    32,
			weight: 50,
		},
		{
			name:   "Dmitry",
			age:    31,
			weight: 80,
		},
	}
	file, err := os.Create("people.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	for _, person := range people {
		fmt.Fprintf(file,
			"%-10s %-10d %-10.3f\n",
			person.name, person.age, person.weight)
	}
}
