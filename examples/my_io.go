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

	file, createError := os.Create("user.txt")
	errorHandler("create_error", createError)
	_, printError := fmt.Fprintf(file, "%s %d\n", user.name, user.age)
	errorHandler("print_error", printError)
	file.Close()

	savedFile, openError := os.Open("user.txt")
	errorHandler("open_error", openError)

	_, scanError := fmt.Fscanf(savedFile, "%s %d\n", &name, &age)
	errorHandler("scan_error", scanError)

	fmt.Println(name, age)
}

func errorHandler(label string, err error) {
	if err != nil {
		fmt.Println(label, err)
		os.Exit(1)
	}
}
