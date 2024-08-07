package dealmanager

import (
	"goRepos/allen_digital_interview/deal"
	"goRepos/allen_digital_interview/user"
)

type IDealManager interface {
	AddDeal(deal deal.Deal)
	EndDeal(dealId string)
	GetDeal(dealId string) deal.Deal
	UpdateDeal(deal deal.Deal)
	AddUser(user user.User)
	AddNewUserDealMapping(userId, dealId string)
	GetUserDealMappingByUserId(userId string) []string
	ClaimDeal(userId, dealId string)
}
