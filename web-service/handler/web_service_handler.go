package handler

import (
	"github.com/google/uuid"
	"sybo/web-service/data"
)

type WebServiceHandler struct {
	storage data.IStorageHandler
}

type Options struct {
	Storage data.IStorageHandler
}

func NewWebServiceHandler(options Options) WebServiceHandler {
	handler := WebServiceHandler{
		storage: options.Storage,
	}
	return handler
}

func (handler WebServiceHandler) CreateUser(req data.NewUserRequest) data.User {
	user := data.User{
		Name: req.Name,
		ID:   uuid.NewString(),
	}

	handler.storage.InsertUser(user)
	return user
}

func (handler WebServiceHandler) SaveState(uid string, state data.State) {
	oldState := handler.LoadState(uid)
	if oldState.Score > state.Score {
		// Keep old highscore if higher
		state.Score = oldState.Score
	}
	// Assuming that client has most up-to-date record of games played
	handler.storage.SetState(uid, state)
}

func (handler WebServiceHandler) LoadState(uid string) data.State {
	return handler.storage.GetState(uid)
}

func (handler WebServiceHandler) UpdateFriends(uid string, friendUids []string) {
	handler.storage.SetFriends(uid, friendUids)
}

func (handler WebServiceHandler) GetFriends(uid string) []string {
	return handler.storage.GetFriends(uid)
}

func (handler WebServiceHandler) GetAllUsers() []data.User {
	return handler.storage.GetAllUsers()
}
