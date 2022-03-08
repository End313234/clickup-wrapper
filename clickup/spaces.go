package clickup

import (
	"fmt"
	"strconv"

	"github.com/End313234/clickup-wrapper/internal/constants"
	"github.com/End313234/clickup-wrapper/internal/http/request"
)

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
		client.Cache.Spaces.Add(space)
		return space, err
	}

	client.Cache.Spaces.Add(space)
	return space, nil
}

// Gets the spaces of a Team
func (client *Client) GetSpaces(teamId string, archived bool) ([]Space, error) {
	var spaces Spaces

	err := request.MakeRequest(request.CustomRequest{
		Method:      "GET",
		URL:         fmt.Sprintf("%s/team/%s/space?archived=%s", constants.BASE_URL, teamId, strconv.FormatBool(archived)),
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
