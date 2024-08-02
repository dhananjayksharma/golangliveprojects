package queries

import (
	"context"
	"golangliveprojects/iplplayers/pkg/v1/responses"

	"gorm.io/gorm"
)

type persistentSQLDBStore struct {
	db *gorm.DB
}

func NewPersistentSQLDBStore(dbconn *gorm.DB) PersistentSQLDBStorer {
	return &persistentSQLDBStore{db: dbconn}
}

type PersistentSQLDBStorer interface {
	// players
	PlayerListQuery(ctx context.Context, playerResponse *[]responses.PlayerResponse) error
	PlayerListQueryMatches(ctx context.Context, playerResponse *[]responses.PlayerResponse) error
	PlayerListQueryPlayerDetails(ctx context.Context, playerResponse *[]responses.PlayerResponse, playerCode string) error
	AddPlayerQuery(ctx context.Context, playerResponse *[]responses.PlayerResponse, requests.PlayerAddRequest) error

	// // matches
	// MatcheListQuery(ctx context.Context, playerResponse *[]responses.PlayerResponse) error
}
