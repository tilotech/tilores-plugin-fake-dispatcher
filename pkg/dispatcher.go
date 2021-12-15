package pkg

import (
	"context"

	"github.com/google/uuid"
	api "github.com/tilotech/tilores-plugin-api"
	"github.com/tilotech/tilores-plugin-api/dispatcher"
)

type FakeDispatcher struct {
	records [10]*api.Record
	index   int
	length  int
}

func (f *FakeDispatcher) Entity(_ context.Context, id string) (*api.Entity, error) {
	return &api.Entity{
		ID:      id,
		Records: f.records[0:f.length],
	}, nil
}

func (f *FakeDispatcher) Submit(_ context.Context, records []*api.Record) (*dispatcher.SubmissionResult, error) {
	for _, record := range records {
		f.addRecord(record)
	}
	return &dispatcher.SubmissionResult{
		RecordsAdded: len(records),
	}, nil
}

// Search finds all matching records and returns a slice of Entity
//
// The fake search will return maximum one entity which includes all matching records, unlike the real search.
// Not all search parameters need to match a record field to consider the record a match, one is enough.
func (f *FakeDispatcher) Search(_ context.Context, parameters *api.SearchParameters) ([]*api.Entity, error) {
	matchingRecords := make([]*api.Record, 0, f.length)
	for i := 0; i < f.length; i++ {
		record := f.records[i]
		for key, value := range *parameters {
			if record.Data[key] == value {
				matchingRecords = append(matchingRecords, record)
				break
			}
		}
	}
	if len(matchingRecords) == 0 {
		return []*api.Entity{}, nil
	}
	return []*api.Entity{{
		ID:      uuid.New().String(),
		Records: matchingRecords,
	}}, nil
}

func (f *FakeDispatcher) addRecord(record *api.Record) {
	f.records[f.index] = record
	f.index++
	if f.index == 10 {
		f.index = 0
	}
	if f.length < 10 {
		f.length++
	}
}
