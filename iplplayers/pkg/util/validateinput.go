package util

import (
	"errors"
	"fmt"
	"golangliveprojects/iplplayers/pkg/constants"
	"golangliveprojects/iplplayers/pkg/customerrors"

	"github.com/google/uuid"
)

func ValidatePlayerCode(playerCode string) error {
	//indpl4538a6bf-7e85-4bd5-b7bc-c6584a2cfbe0
	playerPrefix := playerCode[:constants.CONST_NUMBER_FIVE]
	if playerPrefix != constants.CONST_PLAYER_PREFIX {
		err := fmt.Sprintf(customerrors.ERR_INVALID_PLAYER_CODE, playerCode)
		return errors.New(err)
	}
	return nil
}

func GetNewPlayerCode() string {
	//indpl4538a6bf-7e85-4bd5-b7bc-c6584a2cfbe0
	newID := uuid.New().String()
	// if err != nil {
	// 	err := fmt.Sprintf(customerrors.ERR_PLAYER_CODE_GENERATION_FAILED)
	// 	return "", errors.New(err)
	// }
	return fmt.Sprintf("%s%s", constants.CONST_PLAYER_PREFIX, newID)
}
