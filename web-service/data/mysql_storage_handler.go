package data

type MysqlStorageHandler struct {
    
}

func (s MysqlStorageHandler) InsertUser(user *User) {

}

func (s MysqlStorageHandler) GetUser(id string) User {
    return User{
        Name: "fake",
        ID: "id",
    }
}

func (s MysqlStorageHandler) GetAllUsers() []User {
    return make([]User, 0, 0)
}

func (s MysqlStorageHandler) GetState(id string) State {
    return State{
        GamesPlayed: 42,
        Score: 358,
    }
}

func (s MysqlStorageHandler) SetState(id string, state State) {
    
}

func (s MysqlStorageHandler) SetFriends(id string, friendIds []string) {

}

func (s MysqlStorageHandler) GetFriends(id string) []string {
    return make([]string, 0, 0)
}

