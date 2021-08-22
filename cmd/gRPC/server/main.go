package main

import (
	"github.com/develamit/go-microservice-layout/pkg/fileWrite"
	"fmt"
)

func main() {
    fmt.Printf("Main Function")

    astatus := fileWrite.AStatus{
        Id:          2,
        Name:   "Vihaan1231",
    }

    fileWrite.Writer(astatus)

}
