package queries

import (
	"context"
	"golangliveprojects/iplplayers/pkg/v1/responses"
)

type MockPersistentSQLDBStore struct {
}

var _ PersistentSQLDBStorer = (*MockPersistentSQLDBStore)(nil)

// PlayerListQuery
func (ms *MockPersistentSQLDBStore) PlayerListQuery(ctx context.Context, playerResponse *[]responses.PlayerResponse) error {
	return nil
}

// PlayerListQueryMatches
func (ms *MockPersistentSQLDBStore) PlayerListQueryMatches(ctx context.Context, playerResponse *[]responses.PlayerResponse) error {
	return nil
}

// PlayerListQueryPlayerDetails
func (ms *MockPersistentSQLDBStore) PlayerListQueryPlayerDetails(ctx context.Context, playerResponse *[]responses.PlayerResponse, playerCode string) error {
	return nil
}
