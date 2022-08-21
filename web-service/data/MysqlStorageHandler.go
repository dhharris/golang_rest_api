package data

type MysqlStorageHandler struct {
    
}

func (s MysqlStorageHandler) insertUser(user *User) {

}

func (s MysqlStorageHandler) getUser(id string) User {
    return *new(User)
}

func (s MysqlStorageHandler) getAllUsers() []User {
    return make([]User, 0, 0)
}

func (s MysqlStorageHandler) getState(id string) State {
    return *new(State)
}

func (s MysqlStorageHandler) setState(id string, state State) {
    
}

func (s MysqlStorageHandler) setFriends(id string, friendIds []string) {

}

func (s MysqlStorageHandler) getFriends(id string) []string {
    return make([]string, 0, 0)
}

