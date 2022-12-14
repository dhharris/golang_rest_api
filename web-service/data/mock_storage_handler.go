package data

import "golang.org/x/exp/maps"

type MockStorageHandler struct {
	users   map[string]User
	states  map[string]State
	friends map[string]Friends
}

func NewMockStorageHandler() MockStorageHandler {
	return MockStorageHandler{
		users:   make(map[string]User),
		states:  make(map[string]State),
		friends: make(map[string]Friends),
	}
}

func (s MockStorageHandler) InsertUser(user User) {
	s.users[user.ID] = user
}

func (s MockStorageHandler) GetUser(id string) (User, error) {
	return s.users[id], nil
}

func (s MockStorageHandler) GetAllUsers() []User {
	return maps.Values(s.users)
}

func (s MockStorageHandler) GetState(id string) (State, error) {
	return s.states[id], nil
}

func (s MockStorageHandler) SetState(id string, state State) {
	s.states[id] = state
}

func (s MockStorageHandler) SetFriends(id string, friends Friends) {
	s.friends[id] = friends
}

func (s MockStorageHandler) GetFriends(id string) (Friends, error) {
	return s.friends[id], nil
}
