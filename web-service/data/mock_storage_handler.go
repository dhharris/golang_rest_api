package data

import "golang.org/x/exp/maps"

type MockStorageHandler struct {
	users   map[string]User
	states  map[string]State
	friends map[string][]string
}

func NewMockStorageHandler() MockStorageHandler {
	return MockStorageHandler{
		users:   make(map[string]User),
		states:  make(map[string]State),
		friends: make(map[string][]string),
	}
}

func (s MockStorageHandler) InsertUser(user User) {
	s.users[user.ID] = user
}

func (s MockStorageHandler) GetUser(id string) User {
	return s.users[id]
}

func (s MockStorageHandler) GetAllUsers() []User {
	return maps.Values(s.users)
}

func (s MockStorageHandler) GetState(id string) State {
	return s.states[id]
}

func (s MockStorageHandler) SetState(id string, state State) {
	s.states[id] = state
}

func (s MockStorageHandler) SetFriends(id string, friendIds []string) {
	s.friends[id] = friendIds
}

func (s MockStorageHandler) GetFriends(id string) []string {
	return s.friends[id]
}
