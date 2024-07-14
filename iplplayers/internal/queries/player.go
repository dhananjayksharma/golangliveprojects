package queries

import (
	"context"
	"errors"
	"golangliveprojects/iplplayers/pkg/v1/responses"

	"gorm.io/gorm"
)

func ListQuery(ctx context.Context, playerResponse *[]responses.PlayerResponse, db *gorm.DB) error {

	result := db.WithContext(ctx).Model(&responses.PlayerResponse{}).Select("id, player_code, player_name, player_dob, player_country, player_category, status, created_dt").Scan(&playerResponse)
	var err error
	if result.RowsAffected == 0 {
		err = errors.New("error record not found")
		return err
	}

	return nil
}
func ListQueryMatches(ctx context.Context, playerResponse *[]responses.PlayerResponse, db *gorm.DB) error {

	result := db.WithContext(ctx).Model(&responses.PlayerResponse{}).Select("id, player_code, player_name, player_dob, player_country, player_category, status, created_dt").Scan(&playerResponse)
	var err error
	if result.RowsAffected == 0 {
		err = errors.New("error record not found")
		return err
	}

	return nil
}
func ListQueryPlayerDetails(ctx context.Context, playerResponse *[]responses.PlayerResponse, db *gorm.DB) error {

	result := db.WithContext(ctx).Model(&responses.PlayerResponse{}).Select("id, player_code, player_name, player_dob, player_country, player_category, status, created_dt").Scan(&playerResponse)
	var err error
	if result.RowsAffected == 0 {
		err = errors.New("error record not found")
		return err
	}

	return nil
}
