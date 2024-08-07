package dealmanager

import (
	"fmt"
	"time"

	"goRepos/allen_digital_interview/deal"
)

func (DealManager DealManager) IsValidDeal(userId string, deal deal.Deal) bool {
	if !deal.Active {
		fmt.Println("Deal with id:", deal.GetId(), "is not active anymore")
		return false
	}

	dealIdsInDb := DealManager.GetUserDealMappingByUserId(userId)

	for _, dealIdInDb := range dealIdsInDb {
		if dealIdInDb == deal.GetId() {
			fmt.Println("User can only claim the deal once")
			return false
		}
	}

	if deal.NumberOfItems == 0 {
		fmt.Println("No more items left in the deal")
		return false
	}

	if time.Now().Unix() > deal.ExpiredAt {
		fmt.Println("Deal has expired")
		return false
	}

	return true
}
