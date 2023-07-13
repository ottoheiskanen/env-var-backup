package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type FormattedVar struct {
	Variable string `json:"variable"`
	Value    string `json:"value"`
}

func separateAndPush(sep string) []FormattedVar {
	var f []FormattedVar
	for _, env := range os.Environ() {
		pair := strings.SplitN(env, sep, 2)
		record := FormattedVar{Variable: pair[0], Value: pair[1]}
		f = append(f, record)
	}

	return f
}

func saveToJSON(f *[]FormattedVar, filename string) {
	file, err := os.Create(filename + ".json")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()
	writer := json.NewEncoder(file)

	err = writer.Encode(f)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	formattedVars := separateAndPush("=")
	fmt.Printf("%+v", formattedVars[2])
	saveToJSON(&formattedVars, "env")
}
