package user

func GetIndexByUser(userName string) int {
	var userIndex int
	for index, value := range users {
		if value.name == userName {
			userIndex = index
		}
	}
	return userIndex
}
