package handler 

import (
    "sybo/web-service/data"
    "github.com/google/uuid"
)

type WebServiceHandler struct {
    storage data.IStorageHandler
}

func (handler WebServiceHandler) createUser(name string) data.User {
    user := data.User{
        Name: name,
        ID: uuid.NewString(),
    }
    handler.storage.InsertUser(user)
    return user
}

func (handler WebServiceHandler) saveState(uid string, state data.State) {
    oldState := handler.loadState(uid)
    if oldState.Score > state.Score {
        // Keep old highscore if higher
        state.Score = oldState.Score
    }
    // Assuming that client has most up-to-date record of games played
    handler.storage.SetState(uid, state);
}

func (handler WebServiceHandler) loadState(uid string) data.State {
    return handler.storage.GetState(uid)
}

func (handler WebServiceHandler) updateFriends(uid string, friendUids []string) {
    handler.storage.SetFriends(uid, friendUids) 
}

func (handler WebServiceHandler) getFriends(uid string) []string {
    return handler.storage.GetFriends(uid) 
}

func (handler WebServiceHandler) getAllUsers() []data.User {
    return handler.storage.GetAllUsers()    
}


