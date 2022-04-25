package users

import (
	"fmt"
	"strings"
	"sync"
)

func UserSelectors() {
	var wg sync.WaitGroup
	wg.Add(4)

	go func() {
		defer wg.Done()
		fmt.Println("\ngetAll", USERS.getAll())
	}()

	go func() {
		defer wg.Done()
		fmt.Println("\nfindAllByRules", USERS.findAllByRules(TFindRules{
			car: TCar{
				model: "KIA",
			},
		}))
	}()

	go func() {
		defer wg.Done()
		fmt.Println("\nfindFirstByName", USERS.findFirstByName("Ark"))
	}()

	go func() {
		defer wg.Done()
		fmt.Println("\nfindAllByName", USERS.findAllByName("lf"))
	}()

	wg.Wait()
}

func (users TUsers) getAll() TUsers {
	return users
}

func (users TUsers) findFirstByName(name string) (foundUser TUserData) {
	for _, user := range users {
		if user.firstname == name {
			foundUser = user
			break
		}
	}
	return
}

func (users TUsers) findAllByName(name string) (foundUsers TUsers) {
	for _, user := range users {
		if strings.Contains(user.firstname, name) {
			foundUsers = append(foundUsers, user)
		}
	}
	return
}

type TFindRules TUserData

func (users TUsers) findAllByRules(rules TFindRules) (foundUsers TUsers) {
	for _, user := range users {
		namesAreEqual := rules.firstname == user.firstname
		agesAreEqual := rules.age == user.age
		carModelsAreEqual := rules.car.model == user.car.model
		carCostsAreEqual := rules.car.cost == user.car.cost

		if namesAreEqual || agesAreEqual || carModelsAreEqual || carCostsAreEqual {
			foundUsers = append(foundUsers, user)
		}
	}
	return
}
