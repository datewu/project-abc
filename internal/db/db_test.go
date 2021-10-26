package db

import (
	"context"
	"testing"

	"github.com/datewu/gtea/utils"
)

func TestNewVideo(t *testing.T) {
	if utils.InGithubCI() {
		return
	}
	url := "postgresql://user:password@localhost:5432/dbname?connect_timeout=3&sslmode=disable"
	ctx := context.Background()
	err := initPool(ctx, url)
	if err != nil {
		t.Error(err)
	}
}
