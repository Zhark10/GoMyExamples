package osy

import (
	"os"
	"fmt"
)

func WorkWithFile() {
	file, err := os.Create("hello.txt")     
    if err != nil{                        
        fmt.Println("Unable to create file:", err) 
        os.Exit(1)                        
    }
    defer file.Close()                      
    fmt.Println(file.Name())   
}