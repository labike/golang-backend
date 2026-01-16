package util

import (
	"context"
	"go-admin/common/constant"
	"go-admin/pkg/redis"
	"time"
)

// 存取验证码
var ctx = context.Background()

type RedisStore struct{}

func (r RedisStore) Set(id string, value string) error {
	key := constant.LOGIN_CODE + id
	err := redis.RedisDb.Set(ctx, key, value, time.Minute*5).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r RedisStore) Get(id string, clear bool) string {
	key := constant.LOGIN_CODE + id
	val, err := redis.RedisDb.Get(ctx, key).Result()
	if err != nil {
		return ""
	}
	return val
}

func (r RedisStore) Verify(id string, answer string, clear bool) bool {
	v := RedisStore{}.Get(id, clear)
	return v == answer
}
