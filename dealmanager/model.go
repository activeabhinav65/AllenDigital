package dealmanager

import (
	"goRepos/allen_digital_interview/deal"
	"goRepos/allen_digital_interview/user"
)

type DealManager struct {
	UserMap         map[string]user.User
	DealMap         map[string]deal.Deal
	UserDealMapping map[string][]string
}

func NewDealManager() *DealManager {
	return &DealManager{
		UserMap:         make(map[string]user.User),
		DealMap:         make(map[string]deal.Deal),
		UserDealMapping: make(map[string][]string),
	}
}
