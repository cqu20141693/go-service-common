package redis

import (
	"github.com/cqu20141693/go-service-common/v2/config"
	"log"
	"time"

	"github.com/cqu20141693/go-service-common/v2/event"
	"github.com/go-redis/redis/v8"
)

var RedisDB *redis.Client
var ClusterDB *redis.ClusterClient

func init() {
	event.RegisterHook(event.LogComplete, event.NewHookContext(initRedisDB, "initRedis"))
}

func initRedisDB() {

	sub := config.Sub("cc.redis")
	addr := sub.GetString("addr")
	if addr == "" {
		sentinelNodes := sub.GetStringSlice("sentinel.nodes")
		if sentinelNodes == nil || len(sentinelNodes) == 0 {
			nodes := sub.GetStringSlice("cluster.nodes")
			if nodes == nil || len(nodes) == 0 {
				log.Fatal("redis addr not config")
			} else { //cluster
				ClusterDB = redis.NewClusterClient(configClusterOptions())
			}
		} else { // sentinel
			RedisDB = redis.NewFailoverClient(configSentinelOptions())
		}
	} else { // redis server
		RedisDB = redis.NewClient(configRedisOptions())
	}
}

func configClusterOptions() *redis.ClusterOptions {
	sub := config.Sub("cc.redis")
	options := redis.ClusterOptions{}
	options.Addrs = sub.GetStringSlice("cluster.nodes")
	options.Password = sub.GetString("password")
	options.Username = sub.GetString("username")
	options.DialTimeout = sub.GetDuration("conn-timeout") * time.Second
	options.ReadTimeout = sub.GetDuration("read-timeout") * time.Second
	options.PoolTimeout = sub.GetDuration("pool-timeout") * time.Second
	options.IdleTimeout = sub.GetDuration("idle-timeout") * time.Second
	options.MaxRetries = sub.GetInt("retry")
	options.PoolSize = sub.GetInt("pool-size")
	options.MinIdleConns = sub.GetInt("min-idle-conn")
	return &options
}

func configSentinelOptions() *redis.FailoverOptions {
	sub := config.Sub("cc.redis")
	options := redis.FailoverOptions{}
	options.MasterName = sub.GetString("sentinel.master")
	options.SentinelAddrs = sub.GetStringSlice("sentinel.nodes")
	options.SentinelUsername = sub.GetString("sentinel.username")
	options.SentinelPassword = sub.GetString("sentinel.password")
	options.Password = sub.GetString("password")
	options.DB = sub.GetInt("database")
	options.Username = sub.GetString("username")
	options.DialTimeout = sub.GetDuration("conn-timeout") * time.Second
	options.ReadTimeout = sub.GetDuration("read-timeout") * time.Second
	options.PoolTimeout = sub.GetDuration("pool-timeout") * time.Second
	options.IdleTimeout = sub.GetDuration("idle-timeout") * time.Second
	options.MaxRetries = sub.GetInt("retry")
	options.PoolSize = sub.GetInt("pool-size")
	options.MinIdleConns = sub.GetInt("min-idle-conn")
	return &options
}

func configRedisOptions() *redis.Options {
	sub := config.Sub("cc.redis")
	options := redis.Options{}
	options.Addr = sub.GetString("addr")
	options.Password = sub.GetString("password")
	options.DB = sub.GetInt("database")
	options.Username = sub.GetString("username")
	options.DialTimeout = sub.GetDuration("conn-timeout") * time.Second
	options.ReadTimeout = sub.GetDuration("read-timeout") * time.Second
	options.PoolTimeout = sub.GetDuration("pool-timeout") * time.Second
	options.IdleTimeout = sub.GetDuration("idle-timeout") * time.Second
	options.MaxRetries = sub.GetInt("retry")
	options.PoolSize = sub.GetInt("pool-size")
	options.MinIdleConns = sub.GetInt("min-idle-conn")
	return &options
}
