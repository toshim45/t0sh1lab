package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main(){
	if newUUID, err := uuid.NewV7(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(newUUID)
	}
}
