package handler 

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "sybo/web-service/data"
)

type WebServiceHandler interface {
    *gin.Router
}

func (handler WebServiceHandler) createUser(name string) data.User {

}

func (handler WebServiceHandler) saveState(uid string, state data.State) {
    
}

func (handler WebServiceHandler) loadState(uid string) state.State {
    
}

func (handler WebServiceHandler) updateFriends(uids []string) {
    
}

func (handler WebServiceHandler) getFriends(uid string) []string {
    
}
