package util

import (
	"github.com/nanorand/nanorand"
)

func GenerateActivationCode() (string, error) {
	code, err := nanorand.Gen(6)
	if err != nil {
		return "", err
	}
	return code, nil
}
