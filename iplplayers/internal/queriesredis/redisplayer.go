package queriesredis

import (
	"context"
	"errors"
	"fmt"
	"golangliveprojects/iplplayers/pkg/customerrors"
	"log"
	"strings"
	"time"
)

// GetRegistrationDataByKey --- query all with filter
func (ms *redisCacheDBStore) GetRegistrationDataByKey(ctx context.Context, key string) (code string, err error) {
	status := ms.dbRedis.Get(ctx, key)
	if status.Err() != nil && strings.Contains(customerrors.ERR_REDIS_DB_KEY_NOT_FOUND, status.Err().Error()) {
		log.Printf("error : %s:%s, key:%s", customerrors.ERR_INVALID_PLAYER_ACTIVATION_CODE, customerrors.ERR_REDIS_DB_KEY_NOT_FOUND, key)
		err = errors.New(customerrors.ERR_INVALID_PLAYER_ACTIVATION_CODE)
		return
	}
	code = status.Val()
	return
}

// DeleteRegistrationDataByKey --- delete existing key
func (ms *redisCacheDBStore) DeleteRegistrationDataByKey(ctx context.Context, key string) error {
	status := ms.dbRedis.Del(ctx, key)
	if status.Err() != nil {
		return status.Err()
	}
	return nil
}

// PlayerListQuery --- query all with filter
func (ms *redisCacheDBStore) SaveRegistrationDataByKey(ctx context.Context, keyName string, activateCode string) error {
	var ctxrs = context.Background()
	status := ms.dbRedis.Set(ctxrs, keyName, activateCode, 120*time.Second)
	fmt.Println("status:", status)
	if status.Err() != nil {
		return status.Err()
	}
	return nil
}

//player:activation:code:indpl58a210ea-ab74-483f-bb41-75160e62c030
