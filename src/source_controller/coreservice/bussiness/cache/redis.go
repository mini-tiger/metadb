package cache

import (
	"configcenter/src/common/blog"
	"configcenter/src/storage/dal/redis"
	"context"
	redisv7 "github.com/go-redis/redis/v7"
	"strconv"
	"strings"
	"time"
)

/**
 * @Author: Tao Jun
 * @Since: 2023/1/10
 * @Desc: redis.go
**/

var RedisCli *redisv7.Client
var RedisInter redis.Client

func initRedis(cfg redis.Config) error {
	dbNum, err := strconv.Atoi(cfg.Database)
	if nil != err {
		return err
	}
	if cfg.MaxOpenConns == 0 {
		cfg.MaxOpenConns = 3000
	}

	//var client Client
	if cfg.MasterName == "" {
		option := &redisv7.Options{
			Addr:     cfg.Address,
			Password: cfg.Password,
			DB:       dbNum,
			PoolSize: cfg.MaxOpenConns,
		}
		RedisCli = redisv7.NewClient(option)
	} else {
		hosts := strings.Split(cfg.Address, ",")
		option := &redisv7.FailoverOptions{
			MasterName:       cfg.MasterName,
			SentinelAddrs:    hosts,
			Password:         cfg.Password,
			DB:               dbNum,
			PoolSize:         cfg.MaxOpenConns,
			SentinelPassword: cfg.SentinelPassword,
		}
		RedisCli = redisv7.NewFailoverClient(option)
	}

	err = RedisCli.Ping().Err()
	if err != nil {
		return err
	}

	return err
}

func InitCli(ctx context.Context, rediscfg redis.Config, redisInter redis.Client) {
	topCtx = ctx
	RedisInter = redisInter
	for {
		// util redisCli success
		err := initRedis(rediscfg)
		if err == nil {
			break
		} else {
			blog.Errorf("redisCli init fail:%v %d sec after try", err)
		}
		time.Sleep(time.Second * 30)
	}
}
