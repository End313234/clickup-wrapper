package clickup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpace_Create_Success(t *testing.T) {
	chk := assert.New(t)

	client, err := New(Config{
		Token: CLIENT_TOKEN,
	})
	chk.NoError(err)

	newSpace := Space{
		Name: "new-space",
	}
	newSpace, err = newSpace.Create(client, "31026359")
	chk.NoError(err)
	chk.Equal("new-space", newSpace.Name)
}

func TestSpace_Create_Failure(t *testing.T) {
	chk := assert.New(t)

	client, err := New(Config{
		Token: CLIENT_TOKEN,
	})
	chk.NoError(err)

	newSpace := Space{} // No data is given
	newSpace, err = newSpace.Create(client, "31026359")
	chk.Error(err)
	chk.Exactly(Space{}, newSpace)
}

func TestSpace_Delete_Success(t *testing.T) {
	chk := assert.New(t)

	client, err := New(Config{
		Token: CLIENT_TOKEN,
	})
	chk.NoError(err)

	// Create a new Space do delete
	newSpace := Space{Name: "to_delete"}
	newSpace, err = newSpace.Create(client, "31026359")
	chk.NoError(err)
	chk.Equal("to_delete", newSpace.Name)

	err = newSpace.Delete(client)
	chk.NoError(err)
}

func TestSpace_Delete_Failure(t *testing.T) {
	chk := assert.New(t)

	client, err := New(Config{
		Token: CLIENT_TOKEN,
	})
	chk.NoError(err)

	space := Space{}

	err = space.Delete(client)
	chk.Error(err)
}

func TestSpace_Update_Success(t *testing.T) {
	chk := assert.New(t)

	client, err := New(Config{
		Token: CLIENT_TOKEN,
	})
	chk.NoError(err)

	space, err := client.GetSpace("54904345")
	chk.NoError(err)
	chk.Equal("54904345", space.Id)

	updatedSpace, err := space.Update(Space{
		Name: "updated-test",
	}, client)
	chk.NoError(err)
	chk.Equal("updated-test", updatedSpace.Name)
}

func TestSpace_Update_Failure(t *testing.T) {
	chk := assert.New(t)

	client, err := New(Config{
		Token: CLIENT_TOKEN,
	})
	chk.NoError(err)

	space, err := client.GetSpace("54904345")
	chk.NoError(err)
	chk.Equal("54904345", space.Id)

	updatedSpace, err := space.Update(Space{
		Id: "a",
	}, client) // Try to specify an Id
	chk.Error(err)
	chk.Exactly(Space{}, updatedSpace)
}
