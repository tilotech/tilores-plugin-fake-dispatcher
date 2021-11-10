package dispatcher

import (
	"context"
	"net/rpc"

	api "gitlab.com/tilotech/tilores-plugin-api"
)

type client struct {
	client *rpc.Client
}

func (c *client) Entity(ctx context.Context, id string) (*api.Entity, error) {
	var entity api.Entity
	err := c.client.Call(
		"Plugin.Entity",
		map[string]interface{}{
			"id": id,
		},
		&entity,
	)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}