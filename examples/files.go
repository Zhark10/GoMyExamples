package examples

import (
	"fmt"
	"io/ioutil"
	"os"
)

const filename = "hello.txt"

func WriteTxtFile() {
	str_for_file := []byte("Hello, world!!!")
	write_err := ioutil.WriteFile(filename, str_for_file, 0600)

	if write_err != nil {
		fmt.Println(write_err)
	}

	file_data, read_error := ioutil.ReadFile(filename)

	if read_error != nil {
		fmt.Println(read_error)
	}

	fmt.Println(string(file_data))
	os.Remove(filename)
}
