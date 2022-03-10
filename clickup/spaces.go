package clickup

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/End313234/clickup-wrapper/internal/constants"
	"github.com/End313234/clickup-wrapper/internal/http/request"
)

// Represents a Clickup Space
type Space struct {
	Id                string   `json:"id,omitempty"`
	Name              string   `json:"name,omitempty"`
	Private           bool     `json:"private,omitempty"`
	Statuses          []Status `json:"statuses,omitempty"`
	MultipleAssignees bool     `json:"multiple_assignees,omitempty"`
	Features          Features `json:"features,omitempty"`
	Archived          bool     `json:"archived,omitempty"`
}

// Creates a new Space.
func (space *Space) Create(client *Client, teamId string) (Space, error) {
	var newSpace Space

	data, _ := json.Marshal(space)
	reader := ioutil.NopCloser(strings.NewReader(string(data[:])))
	defer reader.Close()

	err := request.MakeRequest(request.CustomRequest{
		Method:      "POST",
		URL:         fmt.Sprintf("%s/team/%s/space", constants.BASE_URL, teamId),
		AccessToken: client.AccessToken,
		Value:       &newSpace,
		Body:        reader,
	})
	if err != nil {
		return newSpace, err
	}

	client.Cache.Spaces.Add(newSpace)
	return newSpace, nil
}

func (space *Space) Delete(client *Client) error {
	err := request.MakeRequest(request.CustomRequest{
		Method:      "DELETE",
		URL:         fmt.Sprintf("%s/space/%s", constants.BASE_URL, space.Id),
		AccessToken: client.AccessToken,
		Value:       &map[string]string{},
	})
	if err != nil {
		return err
	}

	targetIndex := -1
	for index, element := range *client.Cache.Spaces {
		if element.Id == space.Id {
			targetIndex = index
		}
	}

	client.Cache.Spaces.Remove(targetIndex)
	return nil
}
