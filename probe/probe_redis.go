package probe

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/mittwald/mittnite/config"
)

type redisProbe struct {
	addr     string
	password string
}

func NewRedisProbe(cfg *config.RedisConfig) *redisProbe {
	cfg.URL = resolveEnv(cfg.URL)
	cfg.Password = resolveEnv(cfg.Password)
	cfg.Port = resolveEnv(cfg.Port)

	return &redisProbe{
		addr:     fmt.Sprintf("%s:%s", cfg.URL, cfg.Port),
		password: cfg.Password,
	}
}

func (r *redisProbe) Exec() error {
	client := redis.NewClient(&redis.Options{
		Addr:     r.addr,
		Password: r.password,
	})

	_, err := client.Ping().Result()
	return err
}
