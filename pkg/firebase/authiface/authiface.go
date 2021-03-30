package authiface

import (
	"context"

	firebaseauth "firebase.google.com/go/auth"
)

type Client interface {
	GetUsers(ctx context.Context, identifiers []firebaseauth.UserIdentifier) (*firebaseauth.GetUsersResult, error)
}
