package main

import (
	"fmt"
	"os"

	JPI "github.com/cal1co/jpi/api"
)

const (
	terms = "dictdata/jmdict_english/term_bank_29.json"
)

func main() {
	file, err := os.Open(terms)
	if err != nil {
		fmt.Println(err)
	}
	// byteValue, _ := ioutil.ReadAll(file)
	// json.Unmarshal(byteValue)
	// fmt.Println(len(byteValue))
	defer file.Close()

	JPI.JPI()
}
