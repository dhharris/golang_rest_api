package handler

import (
    "github.com/google/uuid"
    "sybo/web-service/data"
    "testing"
)

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
    storage.InsertUser(data.User{
        Name: "test",
        ID: "uid",
    })
    handler := NewWebServiceHandler(Options{
        Storage: storage,
    })

    want := data.State{
        Score: 420,
        GamesPlayed: 69,
    }

    handler.SaveState("uid", want)

    have := storage.GetState("uid")

    if want != have {
        t.Fatalf(`GetState(%q) = %v, want %v`, "uid", have, want)
    }
}

func TestLoadState(t *testing.T) {

    storage := data.NewMockStorageHandler()
    handler := NewWebServiceHandler(Options{
        Storage: storage,
    })

    want := data.State{
        Score: 420,
        GamesPlayed: 69,
    }
    storage.SetState("uid", want)

    have := handler.LoadState("uid")

    if want != have {
        t.Fatalf(`LoadState(%q) = %v, want %v`, "uid", have, want)
    }
}
