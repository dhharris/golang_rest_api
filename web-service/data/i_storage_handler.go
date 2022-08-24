package data

type IStorageHandler interface {
	InsertUser(user User)
	GetUser(uid string) (User, error)
	GetAllUsers() []User
	GetState(uid string) (State, error)
	SetState(uid string, state State)
	SetFriends(uid string, friends Friends)
	GetFriends(uid string) (Friends, error)
}
