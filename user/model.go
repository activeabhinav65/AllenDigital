package user

type User struct {
	Id      string
	Name    string
	Address string
}

func NewUser(id string, name string, address string) *User {
	return &User{
		Id:      id,
		Name:    name,
		Address: address,
	}
}

func (User User) GetId() string {
	return User.Id
}

func (User User) SetId(id string) {
	User.Id = id
}

func (User User) GetName() string {
	return User.Name
}

func (User User) SetName(name string) {
	User.Name = name
}

func (User User) GetAddress() string {
	return User.Address
}

func (User User) SetAddress(address string) {
	User.Address = address
}
