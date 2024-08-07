package dealmanager

import (
	"fmt"

	"goRepos/allen_digital_interview/deal"
)

func (DealManager DealManager) AddDeal(deal deal.Deal) {
	_, ok := DealManager.DealMap[deal.GetId()]
	if ok {
		fmt.Println("Deal with ID :", deal.GetId(), "already exists in system")
		return
	}

	DealManager.DealMap[deal.GetId()] = deal

	fmt.Println("Deal with ID :", deal.GetId(), "successfully added to system")
}

func (DealManager DealManager) EndDeal(dealId string) {
	_, ok := DealManager.DealMap[dealId]
	if !ok {
		fmt.Println("Deal with ID :", dealId, "does not exist in system")
		return
	}

	dealInDb := DealManager.GetDeal(dealId)
	dealInDb.Active = false

	DealManager.UpdateDeal(dealInDb)

	fmt.Println("Deal with ID :", dealId, "successfully ended")
}

func (DealManager DealManager) GetDeal(dealId string) deal.Deal {
	_, ok := DealManager.DealMap[dealId]
	if !ok {
		fmt.Println("Deal with ID :", dealId, "does not exist in system")
		return deal.Deal{}
	}

	return DealManager.DealMap[dealId]
}

func (DealManager DealManager) UpdateDeal(deal deal.Deal) {
	_, ok := DealManager.DealMap[deal.GetId()]
	if !ok {
		fmt.Println("Deal with ID :", deal.GetId(), "does not exist in system")
		return
	}

	DealManager.DealMap[deal.GetId()] = deal

	fmt.Println("Deal with ID :", deal.GetId(), "successfully updated")
}

func (DealManager DealManager) ClaimDeal(userId, dealId string) {
	_, ok := DealManager.UserMap[userId]
	if !ok {
		fmt.Println("User with ID :", userId, "does not exist in system")
		return
	}

	_, ok = DealManager.DealMap[dealId]
	if !ok {
		fmt.Println("Deal with ID :", dealId, "does not exist in system")
		return
	}

	dealInDb := DealManager.GetDeal(dealId)

	if !DealManager.IsValidDeal(userId, dealInDb) {
		return
	}

	DealManager.AddNewUserDealMapping(userId, dealId)
	dealInDb.NumberOfItems = dealInDb.NumberOfItems - 1
	DealManager.UpdateDeal(dealInDb)

	fmt.Println("Successfully claimed the deal")
}
