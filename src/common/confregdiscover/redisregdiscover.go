package confregdiscover

import (
	"configcenter/src/storage/dal/redis"
	"context"
	"fmt"
)

// ZkRegDiscover config register and discover by zookeeper
type RedisRegDiscover struct {
	rediscli redis.Client
	//cancel         context.CancelFunc
	db      string
	rootCtx context.Context
	//sessionTimeOut time.Duration
}

// NewZkRegDiscover create a object of ZkRegDiscover
func NewRedisRegDiscover(client redis.Client, conf redis.Config) *RedisRegDiscover {
	ctx := context.Background()

	return &RedisRegDiscover{
		rediscli: client,
		rootCtx:  ctx,
		db:       conf.Database,
	}
}

// Ping to ping server
func (RD *RedisRegDiscover) Ping() error {
	return RD.rediscli.Ping(RD.rootCtx).Err()
}

//Write to save config data into zookeeper
func (RD *RedisRegDiscover) Write(path string, data []byte) error {
	return RD.rediscli.Set(RD.rootCtx, path, string(data), 0).Err()
}

func (RD *RedisRegDiscover) Read(path string) (string, error) {

	return RD.rediscli.Get(RD.rootCtx, path).Result()
}

func (RD *RedisRegDiscover) Discover(key string) (<-chan *DiscoverEvent, error) {

	env := make(chan *DiscoverEvent, 1)

	go RD.loopDiscover(RD.rootCtx, key, env)

	return env, nil
}

func (RD *RedisRegDiscover) loopDiscover(discvCtx context.Context, path string, env chan *DiscoverEvent) {
	sub := RD.rediscli.Subscribe(discvCtx, fmt.Sprintf("__keyspace@%s__:%s", RD.db, path))

	for {
		discvEnv := &DiscoverEvent{
			Err: nil,
			Key: path,
		}
		//data, _, watchEnv, err := zkRD.zkcli.GetW(path)
		data, err := RD.Read(path)
		discvEnv.Err = err
		discvEnv.Key = path
		discvEnv.Data = []byte(data)
		// write into discoverEvent channel
		env <- discvEnv

		select {
		case <-discvCtx.Done():
			fmt.Printf("discover path(%s) done\n", path)
			return
		case msg := <-sub.Channel(): // xxx 阻塞
			fmt.Printf("watch found the content of path(%s) changed,msg111:%#v\n", path, msg)
		}
	}
}
