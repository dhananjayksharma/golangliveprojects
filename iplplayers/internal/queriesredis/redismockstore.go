package queriesredis

import (
	"context"
)

type MockRedisCacheDB struct {
}

var _ RedisCacheDBStorer = (*MockRedisCacheDB)(nil)

// GetRegistrationDataByKey
func (ms *MockRedisCacheDB) GetRegistrationDataByKey(ctx context.Context, key string) (string, error) {
	return "", nil
}

// SaveRegistrationDataByKey
func (ms *MockRedisCacheDB) SaveRegistrationDataByKey(ctx context.Context, key string, activateCode string) error {
	return nil
}
