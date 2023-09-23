package main

import (
	"sync"
	"testing"
	"time"

	redsync "github.com/go-redsync/redsync/v4"
	goredis "github.com/go-redsync/redsync/v4/redis/goredis/v9"
	redis "github.com/redis/go-redis/v9"
)

func init() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "root",
	})

	redisPool := goredis.NewPool(redisClient)
	rs := redsync.New(redisPool)

	mu = rs.NewMutex("data-mutex")
}

func TestNormal(t *testing.T) {
	t.Log("[NORMAL]")

	id := 1
	value := 100
	data = map[int]int{1: 100}

	value *= 2
	printData(t, id)
	updateData(id, value)
	printData(t, id)
}

func printData(t *testing.T, id int) {
	if d, e := getData(id); e != nil {
		t.Log("[ERROR]", e)
	} else {
		t.Log("[PRINT]", d)
	}
}

func TestRaceCondition(t *testing.T) {
	id := 1
	maxItr := 10
	data = map[int]int{1: 0}

	var wg sync.WaitGroup

	for i := 1; i <= maxItr; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			incrData(id)
		}()
	}
	time.Sleep(304 * time.Microsecond)
	for i := 1; i <= maxItr; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			printData(t, id)
		}()
	}

	wg.Wait()
}

func TestRaceConditionWithLockEnabled(t *testing.T) {
	EnableLock = true
	id := 1
	maxItr := 10
	data = map[int]int{1: 0}

	var wg sync.WaitGroup

	for i := 1; i <= maxItr; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			incrData(id)
		}()
	}
	time.Sleep(304 * time.Microsecond)
	for i := 1; i <= maxItr; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			printData(t, id)
		}()
	}

	wg.Wait()
}
