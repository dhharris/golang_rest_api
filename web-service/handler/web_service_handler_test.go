package handler

import (
    "github.com/google/uuid"
    "sybo/web-service/data"
    "testing"
)

func createTestHandler() WebServiceHandler {
    options := Options{
        Storage: data.NewMockStorageHandler(),
    }
    handler := NewWebServiceHandler(options)

    return handler
}

func TestCreateUser(t *testing.T) {
    handler := createTestHandler()    

    req := data.NewUserRequest{
        Name: "test",
    }

    user := handler.CreateUser(req)

    if user.Name != "test" {
        t.Fatalf(`CreateUser().Name = %q, want "test"`, user.Name)
    }

    _, err := uuid.Parse(user.ID)

    if err != nil {
        t.Fatalf(`Unable to parse UUID %q, %v, error`, user.ID, err)
    }
}
