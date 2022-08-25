package data

import (
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"reflect"
	"testing"
)

func TestInsertUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	handler := MysqlStorageHandler{
		Driver: db,
	}
	user := User{
		Name: "test",
		ID:   "uuid",
	}

	mock.ExpectExec("INSERT INTO users").WithArgs(user.ID, user.Name).WillReturnResult(sqlmock.NewResult(1, 1))

	handler.InsertUser(user)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}

func TestGetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	handler := MysqlStorageHandler{
		Driver: db,
	}

	user := User{
		Name: "test",
		ID:   "uuid",
	}

	mock.ExpectQuery("SELECT name FROM users").WithArgs(user.ID).WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow(user.Name))

	have, err := handler.GetUser(user.ID)
	want := user

	if err != nil {
		t.Fatal(err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}

	if !reflect.DeepEqual(want, have) {
		t.Errorf(`GetUser(%q) = %v, want %v`, user.ID, have, want)
	}
}

func TestGetAllUsers(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	handler := MysqlStorageHandler{
		Driver: db,
	}

	user := User{
		Name: "test",
		ID:   "uuid",
	}

	mock.ExpectQuery("SELECT uuid, name FROM users").WillReturnRows(sqlmock.NewRows([]string{"uuid", "name"}).AddRow(user.ID, user.Name))

	have := handler.GetAllUsers()
	want := []User{user}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}

	if !reflect.DeepEqual(want, have) {
		t.Errorf(`GetAllUsers() = %v, want %v`, have, want)
	}
}

func TestGetState(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	handler := MysqlStorageHandler{
		Driver: db,
	}

	mock.ExpectQuery("SELECT games_played, score FROM state").WithArgs("uuid").WillReturnRows(sqlmock.NewRows([]string{"games_played", "score"}).AddRow(1, 2))

	have, err := handler.GetState("uuid")
	want := State{
		GamesPlayed: 1,
		Score:       2,
	}

	if err != nil {
		t.Fatal(err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}

	if !reflect.DeepEqual(want, have) {
		t.Errorf(`GetState(%q) = %v, want %v`, "uuid", have, want)
	}
}

func TestSetState(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	handler := MysqlStorageHandler{
		Driver: db,
	}
	state := State{
		GamesPlayed: 1,
		Score:       2,
	}

	mock.ExpectExec("UPDATE state SET games_played").WithArgs(state.GamesPlayed, state.Score, "uuid").WillReturnResult(sqlmock.NewResult(1, 1))

	handler.SetState("uuid", state)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}

func TestSetFriends(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	handler := MysqlStorageHandler{
		Driver: db,
	}

	friends := Friends{
		Friends: []string{"friend1", "friend2"},
	}
	encoded, err := json.Marshal(friends)
	if err != nil {
		t.Fatal(err)
	}

	mock.ExpectExec("UPDATE friends SET friends").WithArgs(encoded, "uuid").WillReturnResult(sqlmock.NewResult(1, 1))

	handler.SetFriends("uuid", friends)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}

func TestGetFriends(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	handler := MysqlStorageHandler{
		Driver: db,
	}

	friends := Friends{
		Friends: []string{"friend1", "friend2"},
	}
	encoded, err := json.Marshal(friends)
	if err != nil {
		t.Fatal(err)
	}

	mock.ExpectQuery("SELECT friends FROM friends").WithArgs("uuid").WillReturnRows(sqlmock.NewRows([]string{"friends"}).AddRow(encoded))

	have, err := handler.GetFriends("uuid")
	want := friends
	if err != nil {
		t.Fatal(err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}

	if !reflect.DeepEqual(want, have) {
		t.Errorf(`GetFriends(%q) = %v, want %v`, "uuid", have, want)
	}
}
