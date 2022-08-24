package data

type IStorageHandler interface {
	InsertUser(user User)
	GetUser(uid string) (User, error)
	GetAllUsers() ([]User, error)
	GetState(uid string) (State, error)
	SetState(uid string, state State)
	SetFriends(uid string, friendUids []string)
	GetFriends(uid string) ([]string, error)
}
