package pkg

import (
	"context"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	api "github.com/tilotech/tilores-plugin-api"
	"github.com/tilotech/tilores-plugin-api/dispatcher"
)

func TestFakeDispatcher(t *testing.T) {
	fixture := &FakeDispatcher{}
	ctx := context.Background()

	parameters := api.SearchParameters{
		"isOdd": true,
	}

	actualSearchResult, err := fixture.Search(ctx, &parameters)
	assert.NoError(t, err)
	assert.Empty(t, actualSearchResult)

	_, err = fixture.Submit(ctx, createSubmitInput(record("1")))
	assert.NoError(t, err)
	actual, err := fixture.Entity(ctx, "foo-id")
	assert.NoError(t, err)
	assert.Equal(t, "foo-id", actual.ID)
	assert.Equal(t, 1, len(actual.Records))
	assert.Equal(t, "1", actual.Records[0].ID)

	actualSearchResult, err = fixture.Search(ctx, &parameters)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(actualSearchResult))
	assert.Equal(t, 1, len(actualSearchResult[0].Records))

	_, err = fixture.Submit(ctx, createSubmitInput(
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
	assert.NoError(t, err)
	actual, err = fixture.Entity(ctx, "foo-id")
	assert.NoError(t, err)
	assert.Equal(t, 10, len(actual.Records))
	assert.Equal(t, "1", actual.Records[0].ID)
	assert.Equal(t, "2", actual.Records[1].ID)
	assert.Equal(t, "10", actual.Records[9].ID)

	actualSearchResult, err = fixture.Search(ctx, &parameters)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(actualSearchResult))
	assert.Equal(t, 5, len(actualSearchResult[0].Records))

	_, err = fixture.Submit(ctx, createSubmitInput(record("11")))
	assert.NoError(t, err)
	assert.Equal(t, 10, len(actual.Records))
	assert.Equal(t, "11", actual.Records[0].ID)
	assert.Equal(t, "2", actual.Records[1].ID)
	assert.Equal(t, "10", actual.Records[9].ID)
}

func record(id string) *api.Record {
	idInt, _ := strconv.Atoi(id)
	return &api.Record{
		ID: id,
		Data: map[string]interface{}{
			"ignoredField": "match",
			"isOdd":        idInt%2 == 1,
		},
	}
}

func createSubmitInput(records ...*api.Record) *dispatcher.SubmitInput {
	return &dispatcher.SubmitInput{Records: records}
}
