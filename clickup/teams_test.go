package clickup

import (
	"testing"

	"github.com/End313234/clickup-wrapper/internal/helpers"
	"github.com/stretchr/testify/assert"
)

var ENV = helpers.GetEnv("../.env")
var CLIENT_TOKEN = ENV["CLIENT_TOKEN"]
var TESTING_WEBHOOK_URL = ENV["TESTING_WEBHOOK_URL"]

func TestCreateWebhook(test *testing.T) {
	chk := assert.New(test)

	client, err := New(Config{
		Token: CLIENT_TOKEN,
	})
	chk.NoError(err)

	team, err := client.FetchTeam("31026359")
	chk.NoError(err)
	chk.Equal("31026359", team.Id)

	// In that case, I want to subscribe to every event
	webhook, err := team.CreateWebhook(client, CreateWebhookBody{
		Endpoint: TESTING_WEBHOOK_URL,
	})
	chk.NoError(err)
	chk.Equal(TESTING_WEBHOOK_URL, webhook.Endpoint)
}

func TestGetWebhooks(test *testing.T) {
	chk := assert.New(test)

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
}
