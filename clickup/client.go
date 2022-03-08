package clickup

import (
	"fmt"

	"github.com/End313234/clickup-wrapper/internal/constants"
	"github.com/End313234/clickup-wrapper/internal/http/request"
)

// Represents a Clickup client
type Client struct {
	AccessToken string
	Cache       Cache
}

// Configuration for the client instance
type Config struct {
	Token string
}

// Makes a new client
func New(config Config) (*Client, error) {
	var currentUser Member

	err := request.MakeRequest(request.CustomRequest{
		Method:      "GET",
		URL:         fmt.Sprintf("%s/user", constants.BASE_URL),
		AccessToken: config.Token,
		Value:       &currentUser,
	})
	if err != nil {
		return &Client{}, err
	}

	return &Client{
		AccessToken: config.Token,
	}, nil
}
