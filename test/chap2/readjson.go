package main

// 解析json文件，遍历
import (
	"encoding/json"
	"fmt"
	"os"
)

const name = "chapter2/sample/data/data.json"

type jsonType struct {
	SITE string `json:"site"`
	LINK string `json:"link"`
	TYPE string `json:"type"`
}

func main() {
	types, _ := openFile()
	for i, v := range types {
		fmt.Println(i, v)
	}
}

func openFile() ([]*jsonType, error) {
	file, e := os.Open(name)
	if e != nil {
		return nil, e
	}
	defer file.Close()
	var jsonty []*jsonType
	e = json.NewDecoder(file).Decode(&jsonty)
	return jsonty, e
}
