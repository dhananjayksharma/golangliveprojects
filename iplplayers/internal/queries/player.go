package queries

import (
	"context"
	"errors"
	"golangliveprojects/iplplayers/pkg/v1/responses"
)

func (ms *persistentSQLDBStore) PlayerListQuery(ctx context.Context, playerResponse *[]responses.PlayerResponse) error {

	result := ms.db.WithContext(ctx).Model(&responses.PlayerResponse{}).Select("id, player_code, player_name, player_dob, player_country, player_category, status, created_dt").Scan(&playerResponse)
	var err error
	if result.RowsAffected == 0 {
		err = errors.New("error record not found")
		return err
	}

	return nil
}
func (ms *persistentSQLDBStore) PlayerListQueryMatches(ctx context.Context, playerResponse *[]responses.PlayerResponse) error {

	result := ms.db.WithContext(ctx).Model(&responses.PlayerResponse{}).Select("id, player_code, player_name, player_dob, player_country, player_category, status, created_dt").Scan(&playerResponse)
	var err error
	if result.RowsAffected == 0 {
		err = errors.New("error record not found")
		return err
	}

	return nil
}
func (ms *persistentSQLDBStore) PlayerListQueryPlayerDetails(ctx context.Context, playerResponse *[]responses.PlayerResponse, playerCode string) error {

	result := ms.db.WithContext(ctx).Debug().Model(&responses.PlayerResponse{}).Select("id, player_code, player_name, player_dob, player_country, player_category, status, created_dt").Where("player_code=?", playerCode).Scan(&playerResponse)
	var err error
	if result.RowsAffected == 0 {
		err = errors.New("error record not found")
		return err
	}

	return nil
}
