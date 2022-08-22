package data

type IStorageHandler interface {
    InsertUser(user User)
    GetUser(uid string) User
    GetAllUsers() []User
    GetState(uid string) State
    SetState(uid string, state State)
    SetFriends(uid string, friendUids []string)
    GetFriends(uid string) []string
}
