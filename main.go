package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func saveToJSON(f []FormattedVar, filename string) {
	bytes, err := json.MarshalIndent(f, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile(filename+".json", bytes, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	formattedVars := separateAndPush("=")
	saveToJSON(formattedVars, "env")
}
