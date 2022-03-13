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
