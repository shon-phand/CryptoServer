package redis

import "github.com/go-redis/redis/v7"

var (
	Client *redis.Client
)

func init() {

	Client = redis.NewClient(&redis.Options{
		Addr:     "redis:6379", // use default Addr
		Password: "",           // no password set
		DB:       0,
	})

	_, err := Client.Ping().Result()
	if err != nil {
		panic("Unable to connect to redis")
	}

}
