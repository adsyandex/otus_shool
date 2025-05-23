package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisLogger struct {
	client *redis.Client
}

func NewRedisLogger(addr string) (*RedisLogger, error) {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return &RedisLogger{client: client}, nil
}

func (l *RedisLogger) LogAction(ctx context.Context, action string, ttl time.Duration) error {
	return l.client.Set(ctx, action, time.Now().String(), ttl).Err()
}

func (l *RedisLogger) Close() error {
	return l.client.Close()
}
