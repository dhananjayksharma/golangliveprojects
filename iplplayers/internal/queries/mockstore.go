package queries

import (
	"context"
	"golangliveprojects/iplplayers/internal/entities"
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

func (ms *MockPersistentSQLDBStore) AddPlayerQuery(ctx context.Context, addRequest *entities.Players) error {
	return nil
}

func (ms *MockPersistentSQLDBStore) UpdatePlayerQuery(ctx context.Context, addRequest *entities.PlayersUpdate, playerCode string) error {
	return nil
}

func (ms *MockPersistentSQLDBStore) GetPlayerByPlayerCode(ctx context.Context, playerResponse *entities.Players, playerCode string) error {
	return nil
}
