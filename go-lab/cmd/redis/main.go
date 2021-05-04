package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/guregu/null.v3"
	"gopkg.in/redis.v5"
)

func main() {
	c := SetupRedis("localhost", "6379")

	//GET SET
	key1 := "key-1"
	c.Set(key1, "val-1", 0)
	val1 := c.Get(key1)
	fmt.Printf("command: %s => %s\r\n", val1, val1.Val())
	key2 := "key-2"
	c.Set(key2, "val-2", 0)
	val2 := c.Get(key2)
	fmt.Printf("command: %s => %s\r\n", val2, val2.Val())

	//MGET MSET
	keys := strings.Split("key-1,key-2", ",")
	vals, err := c.MGet(keys...).Result()
	if err != nil {
		panic(err)
	}
	printVals(vals)

	err = c.MSet("key-3", "val-3", "key-4", "val-4").Err()
	if err != nil {
		fmt.Errorf("err: %v", err)
	}

	//	keys = c.Keys("key-*").Val() //use Scan instead, Keys is blocking
	var cursor uint64
	for {
		keys, cursor, err = c.Scan(0, "key-*", 100).Result()
		if err != nil {
			panic(err)
		}
		if cursor == 0 {
			break
		}
	}

	vals = c.MGet(keys...).Val()
	printVals(vals)

	keys = append(keys, "lock-0")
	keys = append(keys, "key-1")

	vals, err = c.MGet(keys...).Result()
	if err != nil {
		fmt.Errorf("err: %v", err)
	}
	printVals(vals)

	//MGET MSET OBJECT
	mapLen := 3
	models := make(map[string]interface{}, mapLen)
	models["one"] = Item{Name: "one", Amount: 1, Alias: null.StringFrom("siji")}
	models["two"] = Item{Name: "two", Amount: 2, Alias: null.StringFrom("loro")}
	models["three"] = Item{Name: "three", Amount: 3, Alias: null.StringFrom("telu")}
	models["four"] = Item{Name: "40400", Amount: 4, Alias: null.StringFrom("papat")}
	models["five"] = Item{Name: "five", Amount: 5, Alias: null.StringFrom("limo")}

	setMultiple(c, models)

	//	result := make(map[string]interface{}, mapLen)
	result := []Item{}
	keys = []string{"one", "six", "four", "five"}
	getMultiple(c, keys, &result)

	fmt.Printf("result 1 : %v\r\n", result)
	printItem(result)
}

type Item struct {
	Name   string      `json:"item_name" mapstructure:"item_name"`
	Alias  null.String `json:"item_alias" mapstructure:"item_alias"`
	Amount int         `json:"amount"`
}

func printItem(items []Item) {
	for i, r := range items {
		fmt.Printf("result-object %d:%s-%s[%d]\r\n", i, r.Name, r.Alias.String, r.Amount)
	}
}

func getMultiple(c *redis.Client, keys []string, models interface{}) error {
	serializeds, err := c.MGet(keys...).Result()
	if err != nil {
		fmt.Printf("cache retrieval %v %v", keys, err)
	}

	result := make([]interface{}, len(keys))
	for i, serialized := range serializeds {
		str, ok := serialized.(string)
		if !ok {
			continue
		}
		//		var model Item //TODO replace with reflection, still impossible
		var model map[string]interface{}
		err = json.Unmarshal([]byte(str), &model)
		if err != nil {
			panic(err)
		}
		//		models[keys[i]] = model
		result[i] = model
	}
	fmt.Printf("raw-map: %d %v\r\n", len(result), result)
	return mapstructure.Decode(result, models)
}

func setMultiple(c *redis.Client, models map[string]interface{}) error {
	dataLen := len(models)
	data := make([]interface{}, 2*dataLen)
	var i int
	for k, v := range models {
		fmt.Printf("[data] %s %v \r\n", k, v)
		data[i] = k
		i++
		serialized, err := json.Marshal(v)
		if err != nil {
			panic(err)
		}
		data[i] = serialized
		i++
	}
	return c.MSet(data...).Err()
}

//func setMultiple(c *redis.Client, models map[string]interface{}) error {
//	for i, v := range models {
//		err := setSingle(c, i, v)
//		if err != nil {
//			panic(err)
//		}
//	}
//	return nil
//}

func setSingle(c *redis.Client, key string, model interface{}) error {
	serialized, err := json.Marshal(model)
	if err != nil {
		return err
	}
	return c.Set(key, serialized, 24*time.Hour).Err()
}

func printVals(vals []interface{}) {
	fmt.Println("values: ", vals)
	for _, val := range vals {
		if val != nil {
			fmt.Printf("%s\t", val)
		}
	}
	fmt.Println()
}

func SetupRedis(host string, port string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: host + ":" + port,
		DB:   0,
	})
	err := client.Ping().Err()
	if err != nil {
		panic(err)
	}

	return client
}
