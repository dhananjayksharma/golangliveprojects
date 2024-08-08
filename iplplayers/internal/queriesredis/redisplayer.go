package queriesredis

import (
	"context"
	"fmt"
	"time"
)

// GetRegistrationDataByKey --- query all with filter
func (ms *redisCacheDBStore) GetRegistrationDataByKey(ctx context.Context, key string) (string, error) {

	return "123456", nil
}

// PlayerListQuery --- query all with filter
func (ms *redisCacheDBStore) SaveRegistrationDataByKey(ctx context.Context, key string, activateCode string) error {
	var ctxrs = context.Background()
	keyName := fmt.Sprintf("player:activation:code:%s", key)

	status := ms.dbRedis.Set(ctxrs, keyName, activateCode, 120*time.Second)
	fmt.Println("status:", status)
	if status.Err() != nil {
		return status.Err()
	}
	return nil
}

//player:activation:code:indpl58a210ea-ab74-483f-bb41-75160e62c030
