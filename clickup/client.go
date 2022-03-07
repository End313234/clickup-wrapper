package clickup

import (
	"fmt"
	"strconv"
)

type Client struct {
	AccessToken string
	Cache       Cache
}

func (client *Client) GetSpace(spaceId string) (Space, error) {
	var space Space

	space, err := client.Cache.Spaces.Find(func(s Space) bool {
		return s.Id == spaceId
	})
	if err == nil {
		client.Cache.Spaces.Add(space)
		return space, nil
	}

	err = MakeRequest(CustomRequest{
		Method:      "GET",
		URL:         fmt.Sprintf("%s/space/%s", BASE_URL, spaceId),
		AccessToken: client.AccessToken,
		Value:       &space,
	})
	if err != nil {
		client.Cache.Spaces.Add(space)
		return space, err
	}

	client.Cache.Spaces.Add(space)
	return space, nil
}

func (client *Client) GetSpaces(teamId string, archived bool) ([]Space, error) {
	var spaces Spaces

	err := MakeRequest(CustomRequest{
		Method:      "GET",
		URL:         fmt.Sprintf("%s/team/%s/space?archived=%s", BASE_URL, teamId, strconv.FormatBool(archived)),
		AccessToken: client.AccessToken,
		Value:       &spaces,
	})
	if err != nil {
		return []Space{}, err
	}

	for _, space := range spaces.Spaces {
		if _, err := client.Cache.Spaces.Find(func(s Space) bool { return s.Id == space.Id }); err != nil {
			client.Cache.Spaces.Add(space)
		}
	}

	return spaces.Spaces, nil
}

type Config struct {
	Token string
}

func New(config Config) (*Client, error) {
	var currentUser Member

	err := MakeRequest(CustomRequest{
		Method:      "GET",
		URL:         fmt.Sprintf("%s/user", BASE_URL),
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
