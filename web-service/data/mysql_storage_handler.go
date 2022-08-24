package data

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"time"
)

// Used to open a MySQL connection
const kMySql = "mysql"

type MysqlOptions struct {
	// In a production environment, a more sophisticated MySQL library would be
	// used that can deduce the IP address of a given shard
	DbName string
	// DB credentials should be stored securely in encrypted configs.
	// For the purpose of this exercise, the credentials are hard-coded.
	DbUser     string
	DbPassword string
}

func (o MysqlOptions) GetDriver() *sql.DB {
	credentials := fmt.Sprintf("%s:%s@/%s", o.DbUser, o.DbPassword, o.DbName)
	db, err := sql.Open(kMySql, credentials)

	if err != nil {
		log.Fatal(err)
	}

	// Probably should be shorter timeout
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}

type MysqlStorageHandler struct {
	Driver *sql.DB
}

func (s MysqlStorageHandler) InsertUser(user User) {
	query := "INSERT INTO users (uuid, name) VALUES (?, ?)"
	// Safe from SQL injection attacks
	_, err := s.Driver.Exec(query, user.ID, user.Name)

	if err != nil {
		// We never expect errors when creating new users
		log.Fatalf("Error adding user %v to DB: %v", user, err)
	}
}

func (s MysqlStorageHandler) GetUser(id string) (User, error) {
	user := User{ID: id}
	query := "SELECT name FROM users WHERE uuid = ?"
	err := s.Driver.QueryRow(query, id).Scan(&user.Name)

	if err != nil {
		log.Error(err)
	}

	return user, err
}

func (s MysqlStorageHandler) GetAllUsers() []User {
	var users []User

	res, err := s.Driver.Query("SELECT uuid, name FROM users")

	if err != nil {
		log.Error(err)
		return users
	}

	for res.Next() {
		var user User
		// N.B. order must match SELECT
		err := res.Scan(&user.ID, &user.Name)

		if err != nil {
			log.Errorf("Error parsing User info from DB: %v", err)
		}

		users = append(users, user)
	}

	return users
}

func (s MysqlStorageHandler) GetState(id string) (State, error) {
	var state State
	query := "SELECT games_played, score FROM state WHERE uuid = ?"

	err := s.Driver.QueryRow(query, id).Scan(&state.GamesPlayed, &state.Score)

	if err != nil {
		log.Error(err)
	}

	return state, err
}
func (s MysqlStorageHandler) SetState(id string, state State) {
	query := "UPDATE state SET games_played = ?, score = ? WHERE uuid = ?"
	_, err := s.Driver.Exec(query, state.GamesPlayed, state.Score, id)

	if err != nil {
		log.Fatalf("Error updating state for user %q: %v", id, err)
	}
}

func (s MysqlStorageHandler) SetFriends(id string, friendIds []string) {
}

func (s MysqlStorageHandler) GetFriends(id string) ([]string, error) {
	var friends []string
	query := "SELECT friends FROM friends WHERE uuid = ?"

	err := s.Driver.QueryRow(query, id).Scan(&friends)

	if err != nil {
		log.Error(err)
	}

	return friends, err
}
