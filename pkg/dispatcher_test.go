package pkg

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	api "gitlab.com/tilotech/tilores-plugin-api"
)

func TestFakeDispatcher(t *testing.T) {
	fixture := &FakeDispatcher{}
	ctx := context.Background()

	fixture.Submit(ctx, records(record("1")))
	actual, err := fixture.Entity(ctx, "foo-id")
	assert.NoError(t, err)
	assert.Equal(t, "foo-id", actual.ID)
	assert.Equal(t, 1, len(actual.Records))
	assert.Equal(t, "1", actual.Records[0].ID)

	fixture.Submit(ctx, records(
		record("2"),
		record("3"),
		record("4"),
		record("5"),
		record("6"),
		record("7"),
		record("8"),
		record("9"),
		record("10"),
	))
	actual, err = fixture.Entity(ctx, "foo-id")
	assert.NoError(t, err)
	assert.Equal(t, 10, len(actual.Records))
	assert.Equal(t, "1", actual.Records[0].ID)
	assert.Equal(t, "2", actual.Records[1].ID)
	assert.Equal(t, "10", actual.Records[9].ID)

	fixture.Submit(ctx, records(record("11")))
	assert.NoError(t, err)
	assert.Equal(t, 10, len(actual.Records))
	assert.Equal(t, "11", actual.Records[0].ID)
	assert.Equal(t, "2", actual.Records[1].ID)
	assert.Equal(t, "10", actual.Records[9].ID)
}

func record(id string) *api.Record {
	return &api.Record{
		ID: id,
	}
}

func records(records ...*api.Record) []*api.Record {
	return records
}
