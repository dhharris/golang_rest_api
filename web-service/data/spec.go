package data

type NewUserRequest struct {
    Name string `json:"name"`
}

type User struct {
    ID string `json:"id"`
    Name string `json:"name"`
}

type State struct {
    GamesPlayed int `json:"gamesPlayed"`
    Score int `json:"score"`
}

