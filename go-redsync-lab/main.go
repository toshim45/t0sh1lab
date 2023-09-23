package main

import (
	"errors"
	"log"

	"github.com/go-redsync/redsync/v4"
)

var (
	data map[int]int

	ErrNotFound = errors.New("ErrNotFound")

	mu *redsync.Mutex

	EnableLock bool
)

func main() {
	log.Println("[MAIN]")
}
func getData(id int) (int, error) {
	var logPrefix string
	if EnableLock {
		logPrefix = "[LOCKED]"
		if err := mu.Lock(); err != nil {
			log.Fatal(logPrefix+"[LOCK]", err)
		}

		defer func() {
			if ok, err := mu.Unlock(); !ok || err != nil {
				log.Fatal(logPrefix+"[UNLOCK]", err)
			}
		}()
	}

	d, exist := data[id]

	if !exist {
		return 0, ErrNotFound
	}

	return d, nil
}

func incrData(id int) {
	var logPrefix string
	if EnableLock {
		logPrefix = "[LOCKED]"
		if err := mu.Lock(); err != nil {
			log.Fatal(logPrefix+"[LOCK]", err)
		}

		defer func() {
			if ok, err := mu.Unlock(); !ok || err != nil {
				log.Fatal(logPrefix+"[UNLOCK]", err)
			}
		}()
	}

	log.Println(logPrefix + "[START-INCR]")

	value := 10
	d, exist := data[id]

	if exist {
		value += d
	}

	data[id] = value
	log.Println(logPrefix+"[END-INCR]", value)

	return
}

func updateData(id int, value int) error {

	if _, exist := data[id]; !exist {
		return ErrNotFound
	}

	data[id] = value
	log.Println("[UPDATE]", value)
	return nil
}
