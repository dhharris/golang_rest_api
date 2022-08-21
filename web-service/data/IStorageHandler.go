package data

type IStorageHandler interface {
    insertUser()
    getUser() User
    getAllUsers() []User
    getState() State
    setState()
    setFriends()
    getFriends() []string
}
