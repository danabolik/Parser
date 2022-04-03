package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func Parse() {
	path := "examples/files/all-transactions.csv"
	ext := filepath.Ext(path)
	name := filepath.Base(path)
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	jsonO, _ := json.Marshal(bs)
	var result []byte
	json.Unmarshal(jsonO, &result)
	fmt.Println(string(result))
	fileName := strings.TrimSuffix(name, ext)
	fmt.Println(fileName)
	fmt.Println(ext)
}
