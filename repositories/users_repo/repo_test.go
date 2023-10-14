package users_repo

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepo_GET(t *testing.T) {
	ctx := context.Background()
	tc, err := pgx.Connect(context.Background(), "postgres://postgres:vit@localhost:5432/postgres")
	assert.Nil(t, err, err)
	assert.NotNil(t, tc)

	repo := New(ctx, tc)
	assert.NotNil(t, repo)

	usr := repo.GET(1)
	assert.Equal(t, 1, usr.Id)
	assert.Equal(t, "вит", usr.Name)
	assert.Equal(t, "ан", usr.Surname)
}

func TestRepo_Get(t *testing.T) {
}
