package modules

import (
	"context"

	"example.com/m/redis"
)

func DeleteAuth(dUuid string) (int64, error) {

	deleted, err := redis.Client.Del(context.Background(), dUuid).Result()

	if err != nil {
		return 0, err
	}

	return deleted, nil
}
