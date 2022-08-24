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

func (s MockStorageHandler) GetUser(id string) (User, error) {
	return s.users[id], nil
}

func (s MockStorageHandler) GetAllUsers() ([]User, error) {
	return maps.Values(s.users), nil
}

func (s MockStorageHandler) GetState(id string) (State, error) {
	return s.states[id], nil
}

func (s MockStorageHandler) SetState(id string, state State) {
	s.states[id] = state
}

func (s MockStorageHandler) SetFriends(id string, friendIds []string) {
	s.friends[id] = friendIds
}

func (s MockStorageHandler) GetFriends(id string) ([]string, error) {
	return s.friends[id], nil
}
