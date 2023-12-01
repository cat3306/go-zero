package {{.pkgName}}

import (
	"context"
	"errors"
	"fmt"
	"{{.configPkg}}"
	"github.com/go-redis/redis/v8"
)

var (
    {{.exportVar}} *redisClientPool
)

type redisClientPool struct {
	clients map[int]*redis.Client
}
type clientConf struct {
	Options *redis.Options
	DB      []int
}

func newRedisClients(c *clientConf) (*redisClientPool, error)  {
	r := redisClientPool{
		clients: make(map[int]*redis.Client),
	}

	for i := 0; i < len(c.DB); i++ {
		client := redis.NewClient(&redis.Options{
			Network:            c.Options.Network,
			Addr:               c.Options.Addr,
			Dialer:             c.Options.Dialer,
			OnConnect:          c.Options.OnConnect,
			Username:           c.Options.Username,
			Password:           c.Options.Password,
			DB:                 c.DB[i],
			MaxRetries:         c.Options.MaxRetries,
			MinRetryBackoff:    c.Options.MinRetryBackoff,
			MaxRetryBackoff:    c.Options.MaxRetryBackoff,
			DialTimeout:        c.Options.DialTimeout,
			ReadTimeout:        c.Options.ReadTimeout,
			WriteTimeout:       c.Options.WriteTimeout,
			PoolFIFO:           c.Options.PoolFIFO,
			PoolSize:           c.Options.PoolSize,
			MinIdleConns:       c.Options.MinIdleConns,
			MaxConnAge:         c.Options.MaxConnAge,
			PoolTimeout:        c.Options.PoolTimeout,
			IdleTimeout:        c.Options.IdleTimeout,
			IdleCheckFrequency: c.Options.IdleCheckFrequency,
			TLSConfig:          c.Options.TLSConfig,
			Limiter:            c.Options.Limiter,
		})
		if client == nil {
			return nil, errors.New("client nil")
		}
		_, err := client.Ping(context.Background()).Result()
		if err != nil {
			panic(err)
		}
		r.clients[c.DB[i]] = client
	}
	return &r, nil
}
func (r *redisClientPool) Select(args ...int) *redis.Client {
	db := 0
	if len(args) != 0 {
		db = args[0]
	}
	tmp := r.clients[db]
	if tmp == nil {
		panic("invalid db index")
	}
	return tmp
}

func {{.initCode}}() (err error) {
    RedisPool, err = newRedisClients(&clientConf{
		Options: &redis.Options{
			Addr:     fmt.Sprintf("%s:%d", config.AppConf.Redis.Ip, config.AppConf.Redis.Port),
			Password: config.AppConf.Redis.Pwd,
		},
		DB: config.AppConf.Redis.Db,
	})
	return
}