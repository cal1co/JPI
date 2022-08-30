package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	// Server "github.com/cal1co/jpi/server"
)

const (
	terms   = "dictdata/jmdict_english/term_bank_1.json"
	terms2  = "dictdata/jmdict_english/term_bank_2.json"
	terms3  = "dictdata/jmdict_english/term_bank_3.json"
	terms4  = "dictdata/jmdict_english/term_bank_4.json"
	terms5  = "dictdata/jmdict_english/term_bank_5.json"
	terms6  = "dictdata/jmdict_english/term_bank_6.json"
	terms7  = "dictdata/jmdict_english/term_bank_7.json"
	terms8  = "dictdata/jmdict_english/term_bank_8.json"
	terms9  = "dictdata/jmdict_english/term_bank_9.json"
	terms10 = "dictdata/jmdict_english/term_bank_10.json"
	terms11 = "dictdata/jmdict_english/term_bank_11.json"
	terms12 = "dictdata/jmdict_english/term_bank_12.json"
	terms13 = "dictdata/jmdict_english/term_bank_13.json"
	terms14 = "dictdata/jmdict_english/term_bank_14.json"
	terms15 = "dictdata/jmdict_english/term_bank_15.json"
	terms16 = "dictdata/jmdict_english/term_bank_16.json"
	terms17 = "dictdata/jmdict_english/term_bank_17.json"
	terms18 = "dictdata/jmdict_english/term_bank_18.json"
	terms19 = "dictdata/jmdict_english/term_bank_19.json"
	terms20 = "dictdata/jmdict_english/term_bank_20.json"
)

type Entry struct {
	Word      string
	Alternate string
	Form      string
	Blank     string
	Id        int
	Def       []interface{}
	SecondId  string
	Source    string
}
type Entries struct {
	Entries []Entry
}

func main() {

	// dict := levenshteinsearch.CreateDictionary()

	jsonFile, err := os.Open(terms)
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var entries [][]interface{}
	if err := json.Unmarshal(byteValue, &entries); err != nil {
		panic(err)
	}
	var slice []Entry
	for i := 0; i < len(entries); i++ {
		word := fmt.Sprint(entries[i][0])
		alternate := fmt.Sprint(entries[i][1])
		form := fmt.Sprint(entries[i][2])
		blank := fmt.Sprint(entries[i][3])
		id := fmt.Sprint(entries[i][4])
		intId, err := strconv.Atoi(id)
		if err != nil {
			panic(err)
		}
		// def := fmt.Sprint(entries[i][5])
		def, ok := entries[i][5].([]interface{})
		if !ok {
			panic(ok)
		}
		id2 := fmt.Sprint(entries[i][6])
		source := fmt.Sprint(entries[i][7])
		e := Entry{
			Word:      word,
			Alternate: alternate,
			Form:      form,
			Blank:     blank,
			Id:        intId,
			Def:       def,
			SecondId:  id2,
			Source:    source,
		}
		slice = append(slice, e)
		if slice[i].Word == "こんにちわ" {
			fmt.Println(i)
		}
	}
	fmt.Println(slice[1315].Def)
}
