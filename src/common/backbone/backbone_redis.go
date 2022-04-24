package backbone

import (
	"configcenter/src/storage/dal/redis"
	"context"
	"fmt"
	"strings"
)

func newSvcManagerRedisClient(ctx context.Context, svcManagerAddr redis.Config) (redis.Client, error) {

	/*var redisErr error
	if svcManagerAddr.MasterName == "" {
		// MasterName 为空，表示使用直连redis 。 使用Host,Port 做链接redis参数
		service.Session, redisErr = sessions.NewRedisStore(10, "tcp", webSvr.Config.Redis.Address, webSvr.Config.Redis.Password, []byte("secret"))
		if redisErr != nil {
			return fmt.Errorf("failed to create new redis store, error info is %v", redisErr)
		}
	} else {
		// MasterName 不为空，表示使用哨兵模式的redis。MasterName 是Master标记
		address := strings.Split(webSvr.Config.Redis.Address, ";")
		service.Session, redisErr = sessions.NewRedisStoreWithSentinel(address, 10, webSvr.Config.Redis.MasterName, "tcp", webSvr.Config.Redis.Password, []byte("secret"))
		if redisErr != nil {
			return fmt.Errorf("failed to create new redis store, error info is %v", redisErr)
		}
	}*/
	return redis.NewFromConfig(svcManagerAddr)
}

func RedisConfGenerate(RegRedis string) redis.Config {
	var redisConf redis.Config
	redisConfCommlineArrs := strings.Split(RegRedis, ":")
	// xxx 4.21 init redis
	if len(redisConfCommlineArrs) > 1 {
		fmt.Printf("server not is adminserver\n")
		redisConf = redis.Config{
			Address:      fmt.Sprintf("%s:%s", redisConfCommlineArrs[0], redisConfCommlineArrs[1]),
			Password:     redisConfCommlineArrs[3],
			Database:     redisConfCommlineArrs[2],
			MaxOpenConns: 3000,
		}
	}
	return redisConf
}
