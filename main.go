package main

import (
	"fmt"

	"goRepos/allen_digital_interview/deal"
	"goRepos/allen_digital_interview/dealmanager"
	"goRepos/allen_digital_interview/user"
)

func main() {
	dealManager := dealmanager.NewDealManager()

	// add new user1 to the system
	dealManager.AddUser(user.User{
		Id:      "user1",
		Name:    "abhinav",
		Address: "Gurgaon",
	})

	// add already existing user to the system
	dealManager.AddUser(user.User{
		Id:      "user1",
		Name:    "abhinav",
		Address: "Gurgaon",
	})

	// add user2 to the system
	dealManager.AddUser(user.User{
		Id:      "user2",
		Name:    "tanish",
		Address: "Ghaziabad",
	})

	// add user3 to the system
	dealManager.AddUser(user.User{
		Id:      "user3",
		Name:    "rekha",
		Address: "Bombay",
	})

	// add a new deal to the system
	dealManager.AddDeal(deal.Deal{
		Id:            "deal1",
		Name:          "New Year Sale",
		ItemId:        "item1",
		Price:         100,
		NumberOfItems: 10,
		ExpiredAt:     1754584439,
		Active:        true,
	})

	// add already existing deal to the system
	dealManager.AddDeal(deal.Deal{
		Id:            "deal1",
		Name:          "New Year Sale",
		ItemId:        "item1",
		Price:         100,
		NumberOfItems: 10,
		ExpiredAt:     1754584439,
		Active:        true,
	})

	// end a deal which does not exist in system
	dealManager.EndDeal("deal2")

	// end a deal which exists in system
	dealManager.EndDeal("deal1")

	// add a new deal to the system
	dealManager.AddDeal(deal.Deal{
		Id:            "deal2",
		Name:          "Christmas Sale",
		ItemId:        "item1",
		Price:         100,
		NumberOfItems: 10,
		ExpiredAt:     1754584439,
		Active:        true,
	})

	// update deal to increase number of items left
	dealManager.UpdateDeal(deal.Deal{
		Id:            "deal2",
		Name:          "Christmas Sale",
		ItemId:        "item1",
		Price:         100,
		NumberOfItems: 100,
		ExpiredAt:     1754584439,
		Active:        true,
	})

	// claim a deal which is already ended
	dealManager.ClaimDeal("user1", "deal1")

	// claim a deal
	dealManager.ClaimDeal("user1", "deal2")

	// error : cannot claim a deal which is already taken
	dealManager.ClaimDeal("user1", "deal2")

	// claim a deal : user 2
	dealManager.ClaimDeal("user2", "deal2")

	fmt.Println(dealManager.GetDeal("deal2"))

	// add a new deal to the system
	dealManager.AddDeal(deal.Deal{
		Id:            "deal3",
		Name:          "Independence Day Sale",
		ItemId:        "item1",
		Price:         100,
		NumberOfItems: 2,
		ExpiredAt:     1754584439,
		Active:        true,
	})

	// user1 and user2 claims the deal3
	dealManager.ClaimDeal("user1", "deal3")
	dealManager.ClaimDeal("user1", "deal3")

	// user3 cannot claim deal since all items have been sold out
	dealManager.ClaimDeal("user3", "deal3")

	// add a new expired deal to the system
	dealManager.AddDeal(deal.Deal{
		Id:            "deal4",
		Name:          "Independence Day Sale",
		ItemId:        "item1",
		Price:         100,
		NumberOfItems: 2,
		ExpiredAt:     1691426039,
		Active:        true,
	})

	// user3 cannot claim deal since it has expired
	dealManager.ClaimDeal("user3", "deal4")

}
