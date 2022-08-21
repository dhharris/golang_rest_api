package data

import "golang.org/x/exp/maps"

type MockStorageHandler struct {
    users map[string]User
    states map[string]State
    friends map[string][]string
}

func (s MockStorageHandler) insertUser(user User) {
    s.users[user.id] = user 
}

func (s MockStorageHandler) getUser(id string) User {
    return s.users[id]
}

func (s MockStorageHandler) getAllUsers() []User {
    return maps.Values(s.users)    
}

func (s MockStorageHandler) getState(id string) State {
    return s.states[id]
}

func (s MockStorageHandler) setState(id string, state State) {
    s.states[id] = state
}

func (s MockStorageHandler) setFriends(id string, friendIds []string) {
    s.friends[id] = friendIds
}

func (s MockStorageHandler) getFriends(id string) []string {
    return s.friends[id]
}

