package main

import (
	"context"
	"database/sql"
)

type ProcessRecord struct {
	id, time               int
	user, host, command    string
	dbName, state, sqlText sql.NullString
}

type ProcessList struct {
	records []*ProcessRecord
}

type HotSpotRecord struct {
}

type HotSpotList struct {
	records []*HotSpotRecord
}

// Worker thread will call these interface every 2 seconds
type DataSource interface {
	GetProcessList(ctx context.Context) (*ProcessList, error)
	GetHotSpot(ctx context.Context) (*HotSpot, error)
}
