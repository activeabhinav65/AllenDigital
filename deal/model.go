package deal

type Deal struct {
	Id            string
	Name          string
	ItemId        string
	Price         int
	NumberOfItems int
	ExpiredAt     int64
	Active        bool
}

func NewDeal(id string, name string, itemId string, price int, numberOfItems int, expiredAt int64, active bool) *Deal {
	return &Deal{
		Id:            id,
		Name:          name,
		ItemId:        itemId,
		Price:         price,
		NumberOfItems: numberOfItems,
		ExpiredAt:     expiredAt,
		Active:        active,
	}
}

func (Deal Deal) GetId() string {
	return Deal.Id
}

func (Deal Deal) SetId(id string) {
	Deal.Id = id
}
