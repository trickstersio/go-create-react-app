package server

import (
	"fmt"
	"github.com/trickstersio/go-create-react-app/webpack"
)

// User represents current user session
type User struct {
	Email     string
	FirstName string
	LastName  string
}

// ViewData contains data for the view
type ViewData struct {
	CurrentUser  User
	Webpack *webpack.Webpack
}

// NewViewData creates new data for the view
func NewViewData(buildPath string) (ViewData, error) {
	wp, err := webpack.New(buildPath)

	if err != nil {
		return ViewData{}, fmt.Errorf("failed to read webpack configuration: %w", err)
	}

	return ViewData{
		CurrentUser: User{
			Email:     "bill@example.com",
			FirstName: "Bill",
			LastName:  "Black",
		},
		Webpack: wp,
	}, nil
}

