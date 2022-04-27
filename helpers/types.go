package helpers

import "fmt"

type TAny interface{}

func TypesDetection() {
	var point int
	TypeChecker("")
	TypeChecker(0)
	TypeChecker(2)
	TypeChecker(nil)
	TypeChecker(&point)
}

func TypeChecker(valueWithAnyType TAny) {
	switch valueWithAnyType.(type) {
	case int:
		fmt.Println("is int")
	case string:
		fmt.Println("is string")
	case bool:
		fmt.Println("is bool")
	default:
		fmt.Printf("is unknown type %T\n", valueWithAnyType)
	}
}
