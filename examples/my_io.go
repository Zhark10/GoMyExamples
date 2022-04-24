package examples

import (
	"fmt"
	"os"
)

type TUser struct {
	name string
	age  int
}

func MyIo() {
	var name string
	var age int

	user := TUser{
		name: "Ark",
		age:  27,
	}

	file, create_error := os.Create("user.txt")
	errorHandler("create_error", create_error)
	_, print_error := fmt.Fprintf(file, "%s %d\n", user.name, user.age)
	errorHandler("print_error", print_error)
	file.Close()

	saved_file, open_error := os.Open("user.txt")
	errorHandler("open_error", open_error)

	_, scan_error := fmt.Fscanf(saved_file, "%s %d\n", &name, &age)
	errorHandler("scan_error", scan_error)

	fmt.Println(name, age)
}

func errorHandler(label string, err error) {
	if err != nil {
		fmt.Println(label, err)
		os.Exit(1)
	}
}
