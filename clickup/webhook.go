package clickup

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/End313234/clickup-wrapper/internal/constants"
	"github.com/End313234/clickup-wrapper/internal/http/request"
)

// Represents a Clickup Event
type Event string

const (
	TASK_CREATED               Event = "taskCreated"
	TASK_UPDATED               Event = "taskUpdated"
	TASK_DELETED               Event = "taskDeleted"
	TASK_PRIORITY_UPDATE       Event = "taskPriorityUpdated"
	TASK_STATUS_UPDATED        Event = "taskStatusUpdated"
	TASK_ASSIGNEE_UPDATED      Event = "taskAssigneeUpdated"
	TASK_DUE_DATE_UPDATED      Event = "taskDueDateUpdated"
	TASK_TAG_UPDATED           Event = "taskTagUpdated"
	TASK_MOVED                 Event = "taskMoved"
	TASK_COMMENT_POSTED        Event = "taskCommentPosted"
	TASK_COMMENT_UPDATED       Event = "taskCommentUpdated"
	TASK_TIME_ESTIMATE_UPDATED Event = "taskTimeEstimateUpdated"
	TASK_TIME_TRACKED_UPDATED  Event = "taskTimeTrackedUpdated"
	LIST_CREATED               Event = "listCreated"
	LIST_UPDATED               Event = "listUpdated"
	LIST_DELETED               Event = "listDeleted"
	FOLDER_CREATED             Event = "folderCreated"
	FOLDER_UPDATED             Event = "folderUpdated"
	FOLDER_DELETED             Event = "folderDeleted"
	SPACE_CREATED              Event = "spaceCreated"
	SPACE_UPDATED              Event = "spaceUpdated"
	SPACE_DELETED              Event = "spaceDeleted"
	GOAL_CREATED               Event = "goalCreated"
	GOAL_UPDATED               Event = "goalUpdated"
	GOAL_DELETED               Event = "goalDeleted"
	KEY_RESULT_CREATED         Event = "keyResultCreated"
	KEY_RESULT_UPDATED         Event = "keyResultUpdated"
	KEY_RESULT_DELETED         Event = "keyResultDeleted"
)

type WebhookHealth struct {
	Status    string `json:"status,omitempty"`
	FailCount int    `json:"fail_count,omitempty"`
}

// Represents a Clickup Webhook
type Webhook struct {
	Id       string        `json:"id,omitempty"`
	UserId   int           `json:"userid,omitempty"`
	TeamId   int           `json:"team_id,omitempty"`
	Endpoint string        `json:"endpoint,omitempty"`
	ClientId string        `json:"client_id,omitempty"`
	Events   []Event       `json:"events,omitempty"`
	TaskId   string        `json:"task_id,omitempty"`
	ListId   string        `json:"list_id,omitempty"`
	FolderId string        `json:"folder_id,omitempty"`
	SpaceId  string        `json:"space_id,omitempty"`
	Health   WebhookHealth `json:"health,omitempty"`
	Secret   string        `json:"secret,omitempty"`
}

func (wh *Webhook) Delete(client *Client) error {
	err := request.MakeRequest(request.CustomRequest{
		Method:      "DELETE",
		URL:         fmt.Sprintf("%s/webhook/%s", constants.BASE_URL, wh.Id),
		AccessToken: client.AccessToken,
		Value:       &map[string]string{},
	})
	if err != nil {
		return err
	}

	for index, element := range *client.Cache.Webhooks {
		if element.Id == wh.Id {
			client.Cache.Webhooks.Remove(index)
			break
		}
	}

	return nil
}

// Body of a request to update a Webhook
type UpdateWebhookRequest struct {
	Endpoint string  `json:"endpoint,omitempty"`
	Events   []Event `json:"events,omitempty"`
	Status   string  `json:"active,omitempty"`
}

// Body of a response of a request to updated a webhook.
// Pretty much the same thing as `CreateWebhookResponse`
type UpdateWebhookResponse struct {
	Id      string  `json:"id,omitempty"`
	Webhook Webhook `json:"webhook,omitempty"`
}

func (wh *Webhook) Update(client *Client, body UpdateWebhookRequest) (Webhook, error) {
	var webhook UpdateWebhookResponse

	data, _ := json.Marshal(body)
	reader := ioutil.NopCloser(strings.NewReader(string(data[:])))
	defer reader.Close()

	err := request.MakeRequest(request.CustomRequest{
		Method:      "PUT",
		URL:         fmt.Sprintf("%s/webhook/%s", constants.BASE_URL, wh.Id),
		AccessToken: client.AccessToken,
		Value:       &webhook,
		Body:        reader,
	})
	if err != nil {
		return Webhook{}, err
	}

	client.Cache.Webhooks.Update(webhook.Webhook)
	return webhook.Webhook, nil
}
