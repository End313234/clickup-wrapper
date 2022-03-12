package clickup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_New_Success(t *testing.T) {
	chk := assert.New(t)

	client, err := New(Config{
		Token: CLIENT_TOKEN,
	})
	chk.NoError(err)
	chk.Condition(func() (success bool) { return client.AccessToken == CLIENT_TOKEN })
}

func TestClient_New_Failure(t *testing.T) {
	chk := assert.New(t)

	client, err := New(Config{
		Token: "invalid token",
	})
	chk.Error(err)
	chk.Exactly(&Client{}, client)
}

func TestClient_GetSpace_Success(t *testing.T) {
	chk := assert.New(t)

	client, err := New(Config{
		Token: CLIENT_TOKEN,
	})
	chk.NoError(err)

	space, err := client.GetSpace("54904352")
	chk.NoError(err)
	chk.Equal("54904352", space.Id)
}

func TestClient_GetSpace_Failure(t *testing.T) {
	chk := assert.New(t)

	client, err := New(Config{
		Token: CLIENT_TOKEN,
	})
	chk.NoError(err)

	space, err := client.GetSpace("invalid id")
	chk.Error(err)
	chk.Exactly(Space{}, space)
}

func TestClient_GetTeams_Success(t *testing.T) {
	chk := assert.New(t)

	client, err := New(Config{
		Token: CLIENT_TOKEN,
	})
	chk.NoError(err)

	teams, err := client.GetTeams()
	chk.NoError(err)
	chk.Condition(func() (success bool) { return len(teams) > 0 })
}

func TestClient_GetTeams_Failure(t *testing.T) {
	chk := assert.New(t)

	client, _ := New(Config{
		Token: "invalid token",
	}) // Not checking for errors this time since the error should be thrown by `GetTeams`

	teams, err := client.GetTeams()
	chk.Error(err)
	chk.Exactly([]Team{}, teams)
}

func TestClient_GetTeam_Success(t *testing.T) {
	chk := assert.New(t)

	client, err := New(Config{
		Token: CLIENT_TOKEN,
	})
	chk.NoError(err)

	// Populate cache
	teams, err := client.GetTeams()
	chk.NoError(err)
	chk.Condition(func() (success bool) { return len(teams) > 0 })

	team, err := client.GetTeam(teams[0].Id)
	chk.NoError(err)
	chk.Equal(teams[0].Id, team.Id)
}

func TestClient_GetTeam_Failure(t *testing.T) {
	chk := assert.New(t)

	client, err := New(Config{
		Token: CLIENT_TOKEN,
	})
	chk.NoError(err)

	// Populate cache
	teams, err := client.GetTeams()
	chk.NoError(err)
	chk.Condition(func() (success bool) { return len(teams) > 0 })

	team, err := client.GetTeam("invalid id")
	chk.Error(err)
	chk.Exactly(Team{}, team)
}

func TestClient_FetchTeam_Success(t *testing.T) {
	chk := assert.New(t)

	client, err := New(Config{
		Token: CLIENT_TOKEN,
	})
	chk.NoError(err)

	team, err := client.FetchTeam("31026359")
	chk.NoError(err)
	chk.Condition(func() (success bool) { return team.Id == "31026359" })
}

func TestClient_FetchTeam_Failure(t *testing.T) {
	chk := assert.New(t)

	client, err := New(Config{
		Token: CLIENT_TOKEN,
	})
	chk.NoError(err)

	team, err := client.FetchTeam("invalid id")
	chk.Error(err)
	chk.Exactly(Team{}, team)
}
