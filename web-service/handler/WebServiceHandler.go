package handler 

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "sybo/web-service/data"
)

type WebServiceHandler struct {

}

func (handler WebServiceHandler) createUser(name string) data.User {

}

func (handler WebServiceHandler) saveState(uid string, state data.State) {
    
}

func (handler WebServiceHandler) loadState(uid string) data.State {
    
}

func (handler WebServiceHandler) updateFriends(uids []string) {
    
}

func (handler WebServiceHandler) getFriends(uid string) []string {
    
}

func (handler WebServiceHandler) getAllUsers() []data.User {
    
}


