package handler

import (
	"github.com/google/uuid"
	"reflect"
	"sybo/web-service/data"
	"testing"
)

const kTestUID = "uid"

func TestCreateUser(t *testing.T) {
	storage := data.NewMockStorageHandler()
	handler := NewWebServiceHandler(Options{
		Storage: storage,
	})

	req := data.NewUserRequest{
		Name: "test",
	}

	user := handler.CreateUser(req)

	// Return value should be the new user
	if user.Name != "test" {
		t.Fatalf(`CreateUser(%v).Name = %q, want "test"`, req, user.Name)
	}

	_, err := uuid.Parse(user.ID)

	if err != nil {
		t.Fatalf(`Unable to parse UUID %q, %v, error`, user.ID, err)
	}

	// Test data created
	have := storage.GetUser(user.ID)

	if have != user {
		t.Fatalf(`GetUser(%q) = %v, want %v`, user.ID, have, user)
	}
}

func TestSaveState(t *testing.T) {
	storage := data.NewMockStorageHandler()
	handler := NewWebServiceHandler(Options{
		Storage: storage,
	})
	storage.InsertUser(data.User{
		Name: "test",
		ID:   kTestUID,
	})

	want := data.State{
		Score:       420,
		GamesPlayed: 69,
	}

	handler.SaveState(kTestUID, want)

	have := storage.GetState(kTestUID)

	if want != have {
		t.Fatalf(`GetState(%q) = %v, want %v`, kTestUID, have, want)
	}
}

func TestLoadState(t *testing.T) {
	storage := data.NewMockStorageHandler()
	handler := NewWebServiceHandler(Options{
		Storage: storage,
	})

	want := data.State{
		Score:       420,
		GamesPlayed: 69,
	}
	storage.SetState(kTestUID, want)

	have := handler.LoadState(kTestUID)

	if want != have {
		t.Fatalf(`LoadState(%q) = %v, want %v`, kTestUID, have, want)
	}
}

func TestUpdateFriends(t *testing.T) {
	storage := data.NewMockStorageHandler()
	handler := NewWebServiceHandler(Options{
		Storage: storage,
	})

	want := []string{"uid1", "uid2"}

	handler.UpdateFriends(kTestUID, want)

	have := storage.GetFriends(kTestUID)

	if !reflect.DeepEqual(want, have) {
		t.Fatalf(`GetFriends(%q) = %v, want %v`, kTestUID, have, want)
	}

	// Make sure friends can be updated
	want = append(have, "uid3")
	handler.UpdateFriends(kTestUID, want)

	have = storage.GetFriends(kTestUID)

	if !reflect.DeepEqual(want, have) {
		t.Fatalf(`GetFriends(%q) = %v, want %v`, kTestUID, have, want)
	}
}

func TestGetFriends(t *testing.T) {
	storage := data.NewMockStorageHandler()
	handler := NewWebServiceHandler(Options{
		Storage: storage,
	})

	want := []string{"uid1", "uid2"}

	storage.SetFriends(kTestUID, want)

	have := handler.GetFriends(kTestUID)

	if !reflect.DeepEqual(want, have) {
		t.Fatalf(`GetFriends(%q) = %v, want %v`, kTestUID, have, want)
	}
}

func TestGetAllUsers(t *testing.T) {
	storage := data.NewMockStorageHandler()
	handler := NewWebServiceHandler(Options{
		Storage: storage,
	})

	testUser := data.User{
		Name: "test",
		ID:   kTestUID,
	}

	storage.InsertUser(testUser)

	want := storage.GetAllUsers()
	have := handler.GetAllUsers()

	if !reflect.DeepEqual(want, have) {
		t.Fatalf(`GetAllUsers() = %v, want %v`, have, want)
	}
}
