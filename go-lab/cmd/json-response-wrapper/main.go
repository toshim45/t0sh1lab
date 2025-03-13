package main

import (
	"encoding/json"
	"fmt"
)

const (
	/*
		it is common that a response has wrapper model, and the real data is inside the wrapper.
		the model of the data is dynamic so we will use interface for that.
	*/
	rawJson = `
	{
		"status":"ok",
		"data": {"id": 1,"label": "one"},
		"data_list": [
		{"id":2,"label": "two"},
		{"id":3,"label": "three"}
		]
	}
	`
)

type model struct {
	ID    int    `json:"id"`
	Label string `json:"label"`
}

type baseWrapper struct {
	Status string `json:"status"`
}

type wrapper struct {
	baseWrapper
	Data     model   `json:"data"`
	DataList []model `json:"data_list"`
}

func main() {
	var (
		w wrapper
	)

	rawJsonInByte := []byte(rawJson)
	if err := json.Unmarshal(rawJsonInByte, &w); err != nil {
		fmt.Println("Failed Unmarshal", err)
	}

	fmt.Println("[result] status: ", w.Status)
	fmt.Println("[result] data: ", w.Data.ID, w.Data.Label)

	for i, d := range w.DataList {
		fmt.Println("[result] data-list", i, d.ID, d.Label)
	}
}
