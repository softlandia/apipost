package user

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestData_IsEmpty(t *testing.T) {
	usr := NewUser(0, "", "")
	assert.True(t, usr.IsEmpty())

	usr = NewUser(1, "", "")
	assert.False(t, usr.IsEmpty())

	usr = NewUser(0, "", "654654645")
	assert.False(t, usr.IsEmpty())

	usr = NewUser(0, "skdfjkghd", "")
	assert.False(t, usr.IsEmpty())

	usr = NewUser(1, "39485", "sdflksjd")
	assert.False(t, usr.IsEmpty())
}
