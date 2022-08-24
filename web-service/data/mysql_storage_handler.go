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
	query := fmt.Sprintf("INSERT INTO users (uuid, name) VALUES (%q, %q)", user.ID, user.Name)
	_, err := s.Driver.Query(query)

	if err != nil {
		// We never expect errors when creating new users
		log.Fatalf("Error adding user %v to DB: %v", user, err)
	}
}

func (s MysqlStorageHandler) GetUser(id string) (User, error) {
	var user User
	query := fmt.Sprintf("SELECT uuid, name FROM users WHERE uuid = %q", id)
	res, err := s.Driver.Query(query)

	if err != nil {
		log.Error(err)
		return user, err
	}

	err = res.Scan(&user.ID, &user.Name)

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
	query := fmt.Sprintf("SELECT games_played, score FROM state WHERE uuid = %q", id)

	err := s.Driver.QueryRow(query).Scan(&state.GamesPlayed, &state.Score)

	if err != nil {
		log.Error(err)
	}

	return state, err
}
func (s MysqlStorageHandler) SetState(id string, state State) {
	// Check if null - depending on result, insert / update
	var count int
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM state WHERE uuid = %q", id)
	err := s.Driver.QueryRow(countQuery).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	var query string
	if count == 0 {
		query = fmt.Sprintf("INSERT INTO state (uuid, games_played, score) VALUES (%q, %d, %d)", id, state.GamesPlayed, state.Score)
	} else {
		query = fmt.Sprintf("UPDATE state SET games_played = %d, score = %d WHERE uuid = %q", state.GamesPlayed, state.Score, id)
	}
	_, err = s.Driver.Query(query)

	if err != nil {
		log.Fatalf("Error updating state for user %q: %v", id, err)
	}
}

func (s MysqlStorageHandler) SetFriends(id string, friendIds []string) {
}

func (s MysqlStorageHandler) GetFriends(id string) ([]string, error) {
	return make([]string, 0, 0), nil
}
