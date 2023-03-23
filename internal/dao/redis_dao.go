package dao

import (
	"changeme/internal/define"
	"context"
	"github.com/go-redis/redis/v8"
)

var dbPoll = make(map[string]*redis.Client)

func DialDB(connection *define.Connection) error {
	client := redis.NewClient(&redis.Options{
		Addr:     connection.Address + ":" + connection.Port,
		Username: connection.Name,
		Password: connection.Password,
	})
	res, err := client.Info(context.Background(), "keyspace").Result()
	if err != nil {
		return err
	}
	println(res)
	return nil
}
