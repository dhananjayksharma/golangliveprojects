package util

import (
	"fmt"

	"github.com/nanorand/nanorand"
)

func GenerateActivationCode() (string, error) {
	code, err := nanorand.Gen(6)
	if err != nil {
		return "", err
	}
	return code, nil
}

func GetPlayerActicationKey(key string) string {
	keyName := fmt.Sprintf("player:activation:code:%s", key)
	return keyName
}
