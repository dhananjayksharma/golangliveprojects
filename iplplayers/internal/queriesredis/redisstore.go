package queriesredis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type redisCacheDBStore struct {
	dbRedis *redis.Client
}

func NewRedisCacheDBStore(dbconn *redis.Client) RedisCacheDBStorer {
	return &redisCacheDBStore{dbRedis: dbconn}
}

type RedisCacheDBStorer interface {
	GetRegistrationDataByKey(ctx context.Context, key string) (string, error)
	SaveRegistrationDataByKey(ctx context.Context, key string, activateCode string) error
}
