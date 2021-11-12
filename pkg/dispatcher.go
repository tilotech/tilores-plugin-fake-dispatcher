package pkg

import (
	"context"

	api "gitlab.com/tilotech/tilores-plugin-api"
	"gitlab.com/tilotech/tilores-plugin-api/dispatcher"
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
