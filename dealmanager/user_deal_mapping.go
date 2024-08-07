package dealmanager

func (DealManager DealManager) AddNewUserDealMapping(userId, dealId string) {
	DealManager.UserDealMapping[userId] = append(DealManager.UserDealMapping[userId], dealId)
}

func (DealManager DealManager) GetUserDealMappingByUserId(userId string) []string {
	return DealManager.UserDealMapping[userId]
}
