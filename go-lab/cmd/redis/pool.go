package main

import (
	"fmt"

	redigo "github.com/garyburd/redigo/redis"
)

var pool *redigo.Pool

func init() {
	redis_host := "127.0.0.1"
	redis_port := 6379
	pool_size := 20
	pool = redigo.NewPool(func() (redigo.Conn, error) {
		c, err := redigo.Dial("tcp", fmt.Sprintf("%s:%d", redis_host, redis_port))
		if err != nil {
			return nil, err
		}
		return c, nil
	}, pool_size)
}
func Get() redigo.Conn {
	return pool.Get()
}

func main() {
	conn := Get()
	defer conn.Close()
	_, err := conn.Do("SET", "reqid:1", "test-1", "EX", 5)
	if err != nil {
		fmt.Println("error: " + err.Error())
	}
}
