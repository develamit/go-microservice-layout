package main

import (
	"github.com/develamit/go-microservice-layout/pkg/fileWrite"
	//"go-microservice-layout/pkg/fileWrite"
	"fmt"
)

func main() {
    fmt.Printf("Main Function")

    astatus := fileWrite.AStatus{
        Id:          1,
        Name:   "Vihaan"
    }

    fileWrite.writer(astatus)

}
