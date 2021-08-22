package fileWrite

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
)


type AStatus struct {
	Id   int
	Name string
}

func Writer(aData AStatus) {
	fmt.Printf("fileWrite Function")
	file, _ := json.MarshalIndent(aData, "", " ")
	_ = ioutil.WriteFile("/tmp/test.json", file, 0644)
}
