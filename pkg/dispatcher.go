package pkg

import (
	"context"

	api "gitlab.com/tilotech/tilores-plugin-api"
)

type FakeDispatcher struct{}

func (f *FakeDispatcher) Entity(_ context.Context, id string) (*api.Entity, error) {
	// TODO: return entity with last three submitted records
	return &api.Entity{
		ID: id,
	}, nil
}
