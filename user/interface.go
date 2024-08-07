package user

type IUser interface {
	GetId() string
	SetId(id string)
	GetName() string
	SetName(id string)
	GetAddress() string
	SetAddress(id string)
}
