package main

import (
	"context"
	"fmt"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/cainelli/mock-playground/pkg/firebase/authiface"

	"google.golang.org/api/option"
)

func main() {
	if err := Run(); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}

func Run() error {
	app, err := NewFirebase()
	if err != nil {
		return err
	}
	fbauth, err := app.Auth(context.TODO())
	if err != nil {
		return err
	}

	users, err := GetUsers(fbauth)
	if err != nil {
		return err
	}

	fmt.Println(users)

	return nil
}
func GetUsers(app authiface.Client) ([]string, error) {

	res, err := app.GetUsers(context.TODO(), []auth.UserIdentifier{auth.EmailIdentifier{Email: "fernando@cainelli.me"}})
	if err != nil {
		return []string{}, err
	}
	users := []string{}
	for _, u := range res.Users {
		users = append(users, u.Email)
	}
	return users, nil
}

func NewFirebase() (*firebase.App, error) {
	opt := option.WithCredentialsFile("serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}

	return app, err
}
