package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8" // 注意导入的是新版本
	"github.com/spf13/viper"
	"time"
)

var (
	Rdb *redis.Client
)

// 初始化连接
func Init() (err error) {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d",viper.GetString("redis.host"),viper.GetInt("redis.port")),
		Password: viper.GetString("redis.password"),  // no password set
		DB:       viper.GetInt("redis.DB"),   // use default DB
		PoolSize: viper.GetInt("redis.PoolSize"), // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = Rdb.Ping(ctx).Result()
	return err
}
func Close(){
	Rdb.Close()
}
//func main() {
//	ctx := context.Background()
//	if err := initClient(); err != nil {
//		return
//	}
//
//	err := rdb.Set(ctx, "key", "value", 0).Err()
//	if err != nil {
//		panic(err)
//	}
//
//	val, err := rdb.Get(ctx, "key").Result()
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("key", val)
//
//	val2, err := rdb.Get(ctx, "key2").Result()
//	if err == redis.Nil {
//		fmt.Println("key2 does not exist")
//	} else if err != nil {
//		panic(err)
//	} else {
//		fmt.Println("key2", val2)
//	}
//}