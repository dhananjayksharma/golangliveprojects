package queries

import (
	"context"
	"errors"
	"fmt"
	"golangliveprojects/iplplayers/internal/entities"
	"golangliveprojects/iplplayers/pkg/customerrors"
	"golangliveprojects/iplplayers/pkg/util"
	"golangliveprojects/iplplayers/pkg/v1/responses"
	"strings"
)

// PlayerListQuery --- query all with filter
func (ms *persistentSQLDBStore) PlayerListQuery(ctx context.Context, playerResponse *[]responses.PlayerResponse) error {

	result := ms.db.WithContext(ctx).Model(&responses.PlayerResponse{}).Select("id, player_code, player_name, player_dob, player_country, player_category, status, created_dt").Scan(&playerResponse)
	var err error
	if result.RowsAffected == 0 {
		err = errors.New("error record not found")
		return err
	}

	return nil
}

// PlayerListQueryMatches --- will discuss later
func (ms *persistentSQLDBStore) PlayerListQueryMatches(ctx context.Context, playerResponse *[]responses.PlayerResponse) error {

	result := ms.db.WithContext(ctx).Model(&responses.PlayerResponse{}).Select("id, player_code, player_name, player_dob, player_country, player_category, status, created_dt").Scan(&playerResponse)
	var err error
	if result.RowsAffected == 0 {
		err = errors.New("error record not found")
		return err
	}

	return nil
}

// PlayerListQueryPlayerDetails -- by playercode
func (ms *persistentSQLDBStore) PlayerListQueryPlayerDetails(ctx context.Context, playerResponse *[]responses.PlayerResponse, playerCode string) error {

	result := ms.db.WithContext(ctx).Debug().Model(&responses.PlayerResponse{}).Select("id, player_code, player_name, player_dob, player_country, player_category, status, created_dt").Where("player_code=?", playerCode).Scan(&playerResponse)
	var err error
	if result.RowsAffected == 0 {
		err = errors.New("error record not found")
		return err
	}

	return nil
}

func (ms *persistentSQLDBStore) AddPlayerQuery(ctx context.Context, addRequest *entities.Players) error {
	result := ms.db.WithContext(ctx).Create(&addRequest)
	err := result.Error
	if err != nil {
		if strings.Contains(err.Error(), "consts.ErrDuplicateEntry") {
			_errMsg := fmt.Sprintf("consts.ErrAccountAlreadyExists: %s", addRequest.PlayerName)
			return &util.BadRequest{ErrMessage: _errMsg}
		} else {
			return &util.InternalServer{ErrMessage: err.Error()}
		}
	}

	return nil
}

func (ms *persistentSQLDBStore) UpdatePlayerQuery(ctx context.Context, updateRequest *entities.PlayersUpdate, playerCode string) error {
	// result := ms.db.WithContext(ctx).Model(&entities.PlayersUpdate{}).Where("player_code=%s", playerCode).
	result := ms.db.WithContext(ctx).Debug().Model(&entities.PlayersUpdate{}).Where("player_code=?", playerCode).Omit("player_code").Updates(&updateRequest)
	// db.Model(&User{}).Where("active = ?", true).Update("name", "hello")
	err := result.Error
	if err != nil {
		if strings.Contains(err.Error(), "consts.ErrDuplicateEntry") {
			_errMsg := fmt.Sprintf("consts.ErrAccountAlreadyExists: %s", playerCode)
			return &util.BadRequest{ErrMessage: _errMsg}
		} else if strings.Contains(err.Error(), customerrors.ERR_MYSQL_DB_UNKNOWN_COLUMN) {
			return &util.InternalServer{ErrMessage: customerrors.ERR_INTERNAL_SERVER_ERROR}
		} else {
			return &util.InternalServer{ErrMessage: err.Error()}
		}
	}

	return nil
}

// GetPlayerByPlayerCode -- by playercode
func (ms *persistentSQLDBStore) GetPlayerByPlayerCode(ctx context.Context, playerResponse *entities.Players, playerCode string) error {

	result := ms.db.WithContext(ctx).Debug().Model(&entities.Players{}).Select("id, player_code, player_name, player_dob, player_country, player_category, status, created_dt, updated_dt").Where("player_code=?", playerCode).Scan(&playerResponse)
	var err error
	if result.RowsAffected == 0 {
		err = errors.New("error record not found")
		return err
	}

	return nil
}
