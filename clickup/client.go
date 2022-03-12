package clickup

import (
	"errors"
	"fmt"

	"github.com/End313234/clickup-wrapper/internal/constants"
	"github.com/End313234/clickup-wrapper/internal/http/request"
)

// Represents a Clickup client
type Client struct {
	AccessToken string
	Cache       Cache
}

// Represents the response body of a get request to /team
type GetTeamsResponse struct {
	Teams []Team `json:"teams,omitempty"`
}

// Gets a Space by its ID
func (client *Client) GetSpace(spaceId string) (Space, error) {
	var space Space

	space, err := client.Cache.Spaces.Find(func(s Space) bool {
		return s.Id == spaceId
	})
	if err == nil {
		client.Cache.Spaces.Add(space)
		return space, nil
	}

	err = request.MakeRequest(request.CustomRequest{
		Method:      "GET",
		URL:         fmt.Sprintf("%s/space/%s", constants.BASE_URL, spaceId),
		AccessToken: client.AccessToken,
		Value:       &space,
	})
	if err != nil {
		return space, err
	}

	client.Cache.Spaces.Add(space)
	return space, nil
}

// Gets a Team from the cache by its Id.
// It's a shortcut to
//    team, err := client.Cache.Teams.Find(func (t Team) bool { return t.Id == teamId })
//    if err != nil {
//        log.Fatal(err)
//    }
func (client *Client) GetTeam(teamId string) (Team, error) {
	team, err := client.Cache.Teams.Find(func(t Team) bool { return t.Id == teamId })
	if err != nil {
		return Team{}, errors.New("team not found")
	}

	return team, nil
}

// Get all the teams that the client owns to.
// It's preferable to use `client.Cache.Teams`, since
// it won't make a new request to the API. This function
// is planned to be used only if you need to make a request
func (client *Client) GetTeams() ([]Team, error) {
	var teams GetTeamsResponse

	err := request.MakeRequest(request.CustomRequest{
		Method:      "GET",
		URL:         fmt.Sprintf("%s/team", constants.BASE_URL),
		AccessToken: client.AccessToken,
		Value:       &teams,
	})
	if err != nil {
		return teams.Teams, err
	}

	client.Cache.Teams = (*TeamCache)(&teams.Teams)
	return teams.Teams, nil
}

// Fetches a team by making a request to the API.
func (client *Client) FetchTeam(teamId string) (Team, error) {
	var team Team

	teams, err := client.GetTeams()
	if err != nil {
		return team, err
	}

	parsedTeams := (*TeamCache)(&teams)
	team, err = parsedTeams.Find(func(t Team) bool { return t.Id == teamId })
	if err != nil {
		return team, err
	}

	return team, nil
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
		Cache: Cache{
			Users:    &UserCache{},
			Spaces:   &SpaceCache{},
			Webhooks: &WebhookCache{},
			Teams:    &TeamCache{},
		},
	}, nil
}
