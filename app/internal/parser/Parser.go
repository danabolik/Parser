package parser

import (
	"Parser/internal/factories"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Parse() {
	path := "examples/files/all-transactions.csv"
	ext := filepath.Ext(path)
	name := filepath.Base(path)
	fileName := strings.TrimSuffix(name, ext)

	file, err := factories.CreateByExtension(ext)
	if err != nil {
		return
	}

	file.SetFileName(fileName)
	file.SetExtension(ext)

	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}
	r := bufio.NewReader(f)

	var line []byte
	for {
		line, _, err = r.ReadLine()
		if err != nil {
			return
		}

		fmt.Println(string(line))
	}

	//ext := filepath.Ext(path)
	//name := filepath.Base(path)
	//bs, err := ioutil.ReadFile(path)
	//if err != nil {
	//	return
	//}
	//for _, row := range bs {
	//	fmt.Println(row)
	//}
	//
	//jsonO, _ := json.Marshal(bs)
	//var result []byte
	//json.Unmarshal(jsonO, &result)
	//fmt.Println(string(result))
	//fileName := strings.TrimSuffix(name, ext)
	//fmt.Println(fileName)
	//fmt.Println(ext)
}
