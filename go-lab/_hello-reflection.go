package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type AppleItem struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

func main() {
	var num1 int = 3
	tellMe(num1)
	var num2 float64 = 45.34
	tellMe(num2)
	var num3 AppleItem = AppleItem{}
	tellMe(num3)
	// json
	items := make([]interface{}, 2)
	tellMeJson(items, AppleItem{})
	fmt.Println(">>>>> ", items)

	for _, item := range items {
		if valItem, ok := item.(AppleItem); ok {
			fmt.Printf("[result-object] %v %s %d\r\n", valItem, valItem.Name, valItem.Amount)
		}
	}
}

func tellMe(val interface{}) {
	reflection := reflect.ValueOf(val)
	fmt.Printf("%s %s\r\n", reflection.Type().String(), reflection.Kind().String())
}

func tellMeJson(val []interface{}, valType interface{}) {
	var raws []string = []string{`{"name":"mac","amount":10}`, `{"name":"iphone","amount":20}`}
	for i, raw := range raws {
		fmt.Println(">> ", raw)
		//		var model AppleItem
		//		model := reflect.Indirect(reflect.ValueOf(valType))
		model := reflect.New(reflect.TypeOf(valType))
		// model := reflect.New(reflect.TypeOf(valType).Elem())
		err := json.Unmarshal([]byte(raw), model.Interface())
		if err != nil {
			panic(err)
		}
		fmt.Println(">>> ", model)
		val[i] = model
	}
	fmt.Println(">>>> ", len(val), " ", val)
}
