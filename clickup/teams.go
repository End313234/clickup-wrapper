package clickup

import (
	"fmt"
	"strconv"

	"github.com/End313234/clickup-wrapper/internal/constants"
	"github.com/End313234/clickup-wrapper/internal/http/request"
)

// Represents a Clickup Team
type Team struct {
	Id      string   `json:"id,omitempty"`
	Name    string   `json:"name,omitempty"`
	Color   string   `json:"color,omitempty"`
	Avatar  string   `json:"avatar,omitempty"`
	Members []Member `json:"members,omitempty"`
}

// Get the Spaces of a Team
func (team *Team) GetSpaces(client *Client, archived bool) ([]Space, error) {
	var spaces Spaces

	err := request.MakeRequest(request.CustomRequest{
		Method:      "GET",
		URL:         fmt.Sprintf("%s/team/%s/space?archived=%s", constants.BASE_URL, team.Id, strconv.FormatBool(archived)),
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
