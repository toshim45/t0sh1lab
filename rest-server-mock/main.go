package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("path: %s %s\r\n", r.URL.Path[1:], r.Method)
	debugHeader(r)
	if r.Method == "POST" {
		debugBody(r)
	}

	response, err := ioutil.ReadFile("response.json")
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Printf("reply: %s\r\n", string(response))
	fmt.Fprint(w, string(response))
}

func debugHeader(r *http.Request) {
	fmt.Printf("header:\r\n")
	for k, v := range r.Header {
		fmt.Printf("[%s]->%s\r\n", k, v)
	}
}

func debugBody(r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("body: %s\r\n", string(body))
}

func main() {
	port := 8090
	http.HandleFunc("/", handler)
	fmt.Printf("serving port %d\r\n", port)
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
