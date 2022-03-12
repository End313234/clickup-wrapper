package clickup

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/End313234/clickup-wrapper/internal/constants"
	"github.com/End313234/clickup-wrapper/internal/http/request"
)

// Improve that when generics get avaliable
var all []Event = []Event{
	FOLDER_CREATED,
	FOLDER_DELETED,
	FOLDER_UPDATED,
	GOAL_CREATED,
	GOAL_DELETED,
	GOAL_UPDATED,
	KEY_RESULT_CREATED,
	KEY_RESULT_DELETED,
	KEY_RESULT_UPDATED,
	LIST_CREATED,
	LIST_DELETED,
	LIST_UPDATED,
	SPACE_CREATED,
	SPACE_DELETED,
	SPACE_UPDATED,
	TASK_ASSIGNEE_UPDATED,
	TASK_COMMENT_POSTED,
	TASK_COMMENT_UPDATED,
	TASK_COMMENT_POSTED,
	TASK_COMMENT_UPDATED,
	TASK_CREATED,
	TASK_DELETED,
	TASK_DUE_DATE_UPDATED,
	TASK_MOVED,
	TASK_PRIORITY_UPDATE,
	TASK_STATUS_UPDATED,
	TASK_TAG_UPDATED,
	TASK_TIME_ESTIMATE_UPDATED,
	TASK_TIME_TRACKED_UPDATED,
}

// Represents a Clickup Team
type Team struct {
	Id      string   `json:"id,omitempty"`
	Name    string   `json:"name,omitempty"`
	Color   string   `json:"color,omitempty"`
	Avatar  string   `json:"avatar,omitempty"`
	Members []Member `json:"members,omitempty"`
}

// Represents the response body of a request to get the webhooks
type GetWebhookResponse struct {
	Webhooks []Webhook `json:"webhooks,omitempty"`
}

// Represents the request body of a request to create a webhook
type CreateWebhookBody struct {
	Endpoint string  `json:"endpoint,omitempty"`
	Events   []Event `json:"events,omitempty"`
	SpaceId  string  `json:"space_id,omitempty"`
	FolderId string  `json:"folder_id,omitempty"`
	ListId   string  `json:"list_id,omitempty"`
	TaskId   string  `json:"task_id,omitempty"`
}

// Represents the response body of a request to create a webhook
type CreateWebhookResponse struct {
	Id      string  `json:"id,omitempty"`
	Webhook Webhook `json:"webhook,omitempty"`
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

// Created a webhook. Let `events` parameter empty to subscribe every webhook.
func (team *Team) CreateWebhook(client *Client, body CreateWebhookBody, events ...Event) (Webhook, error) {
	var response CreateWebhookResponse

	if len(events) == 0 {
		body.Events = all
	}

	data, _ := json.Marshal(body)
	reader := ioutil.NopCloser(strings.NewReader(string(data[:])))
	defer reader.Close()

	err := request.MakeRequest(request.CustomRequest{
		Method:      "POST",
		URL:         fmt.Sprintf("%s/team/%s/webhook", constants.BASE_URL, team.Id),
		AccessToken: client.AccessToken,
		Value:       &response,
		Body:        reader,
	})
	if err != nil {
		return Webhook{}, err
	}

	client.Cache.Webhooks.Add(response.Webhook)
	return response.Webhook, nil
}

func (team *Team) GetWebhooks(client *Client) ([]Webhook, error) {
	var webhooks GetWebhookResponse

	err := request.MakeRequest(request.CustomRequest{
		Method:      "GET",
		URL:         fmt.Sprintf("%s/team/%s/webhook", constants.BASE_URL, team.Id),
		AccessToken: client.AccessToken,
		Value:       &webhooks,
	})
	if err != nil {
		return []Webhook{}, err
	}

	for _, webhook := range webhooks.Webhooks {
		client.Cache.Webhooks.Add(webhook)
	}

	return webhooks.Webhooks, nil
}
