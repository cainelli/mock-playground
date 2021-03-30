package main_test

import (
	"context"
	"testing"

	firebaseauth "firebase.google.com/go/auth"
	playground "github.com/cainelli/mock-playground/cmd/simple"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

type AuthMock struct {
}

func (am *AuthMock) GetUsers(ctx context.Context, identifiers []firebaseauth.UserIdentifier) (*firebaseauth.GetUsersResult, error) {
	return &firebaseauth.GetUsersResult{
		Users: []*firebaseauth.UserRecord{{UserInfo: &firebaseauth.UserInfo{Email: "fernando@cainelli.me"}}},
	}, nil
}

func TestGetUser(t *testing.T) {
	app := &AuthMock{}
	s, err := playground.GetUsers(app)
	if err != nil {
		t.Error(err)
	}

	assert.Assert(t, is.Contains(s, "fernando@cainelli.me"))
}
