package dealmanager

import (
	"fmt"

	"goRepos/allen_digital_interview/user"
)

func (DealManager DealManager) AddUser(user user.User) {
	_, ok := DealManager.UserMap[user.GetId()]
	if ok {
		fmt.Println("User with ID :", user.GetId(), "already exists in system")
		return
	}

	DealManager.UserMap[user.GetId()] = user

	fmt.Println("User with ID :", user.GetId(), "successfully added to system")
}
