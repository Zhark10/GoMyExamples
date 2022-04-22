package user

import "fmt"

type TUser struct {
	name string
	age  int16
}

var users = [...]TUser{
	{name: "Andrew", age: 26},
	{name: "Rafa", age: 26},
	{name: "Ark", age: 27},
}

func showUserName(userToShow string) {
	fmt.Println(userToShow)
}

func getUser(getIndexByUsername func(string) int, userName string) (returnedUser TUser) {
	defer showUserName(userName)
	defer fmt.Println("Finish")
	index := getIndexByUsername(userName)
	returnedUser = users[index]
	return
}

func ShowUser() {
	result := getUser(GetIndexByUser, "Andrew")
	fmt.Println(result)
}
