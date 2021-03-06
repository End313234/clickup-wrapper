package clickup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Add test case (Failure)
func TestWebhook_Delete_Success(t *testing.T) {
	chk := assert.New(t)

	client, err := New(Config{
		Token: CLIENT_TOKEN,
	})
	chk.NoError(err)

	team, err := client.FetchTeam("31026359")
	chk.NoError(err)
	chk.Equal("31026359", team.Id)

	webhooks, err := team.GetWebhooks(client)
	chk.NoError(err)
	chk.Condition(func() (success bool) { return len(webhooks) > 0 })

	beforeDeleting := webhooks

	// Delete last webhook
	err = webhooks[len(webhooks)-1].Delete(client)
	chk.NoError(err)

	afterDeleting, err := team.GetWebhooks(client)
	chk.NoError(err)
	chk.Condition(func() (success bool) { return len(beforeDeleting) > len(afterDeleting) })
}

func TestWebhook_Delete_Failure(t *testing.T) {
	chk := assert.New(t)

	client, err := New(Config{
		Token: CLIENT_TOKEN,
	})
	chk.NoError(err)

	team, err := client.FetchTeam("31026359")
	chk.NoError(err)
	chk.Equal("31026359", team.Id)

	webhooks := []Webhook{
		{}, // Empty webhook
	}

	// Delete last webhook
	err = webhooks[len(webhooks)-1].Delete(client)
	chk.Error(err)
}

func TestWebhook_Update_Success(t *testing.T) {
	chk := assert.New(t)

	client, err := New(Config{
		Token: CLIENT_TOKEN,
	})
	chk.NoError(err)

	team, err := client.FetchTeam("31026359")
	chk.NoError(err)
	chk.Equal("31026359", team.Id)

	webhook, err := team.GetWebhooks(client)
	chk.NoError(err)
	chk.Condition(func() (success bool) { return len(webhook) > 0 })

	updatedWebhook, err := webhook[len(webhook)-1].Update(client, UpdateWebhookRequest{
		Events: []Event{TASK_CREATED, TASK_DELETED},
		Status: "active",
	})
	chk.NoError(err)
	chk.Equal("active", updatedWebhook.Health.Status)
}

func TestWebhook_Update_Failure(t *testing.T) {
	chk := assert.New(t)

	client, err := New(Config{
		Token: CLIENT_TOKEN,
	})
	chk.NoError(err)

	team, err := client.FetchTeam("31026359")
	chk.NoError(err)
	chk.Equal("31026359", team.Id)

	webhook, err := team.GetWebhooks(client)
	chk.NoError(err)
	chk.Condition(func() (success bool) { return len(webhook) > 0 })

	// Update last webhook
	updatedWebhook, err := webhook[len(webhook)-1].Update(client, UpdateWebhookRequest{
		Events: []Event{"a"},
		Status: "active",
	})
	chk.Error(err)
	chk.Exactly(Webhook{}, updatedWebhook)
}
