package examples

import (
	"fmt"
	"io/ioutil"
	"os"
)

const filename = "hello.txt"

func WriteTxtFile() {
	strForFile := []byte("Hello, world!!!")
	writeErr := ioutil.WriteFile(filename, strForFile, 0600)

	if writeErr != nil {
		fmt.Println(writeErr)
	}

	fileData, readError := ioutil.ReadFile(filename)

	if readError != nil {
		fmt.Println(readError)
	}

	fmt.Println(string(fileData))
	os.Remove(filename)
}
