package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	// Server "github.com/cal1co/jpi/server"
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
	data []Entry
}

func main() {
	dicts := []string{"dictdata/jmdict_english/term_bank_1.json",
		"dictdata/jmdict_english/term_bank_2.json",
		"dictdata/jmdict_english/term_bank_3.json",
		"dictdata/jmdict_english/term_bank_4.json",
		"dictdata/jmdict_english/term_bank_5.json",
		"dictdata/jmdict_english/term_bank_6.json",
		"dictdata/jmdict_english/term_bank_7.json",
		"dictdata/jmdict_english/term_bank_8.json",
		"dictdata/jmdict_english/term_bank_9.json",
		"dictdata/jmdict_english/term_bank_10.json",
		"dictdata/jmdict_english/term_bank_11.json",
		"dictdata/jmdict_english/term_bank_12.json",
		"dictdata/jmdict_english/term_bank_13.json",
		"dictdata/jmdict_english/term_bank_14.json",
		"dictdata/jmdict_english/term_bank_15.json",
		"dictdata/jmdict_english/term_bank_16.json",
		"dictdata/jmdict_english/term_bank_17.json",
		"dictdata/jmdict_english/term_bank_18.json",
		"dictdata/jmdict_english/term_bank_19.json",
		"dictdata/jmdict_english/term_bank_20.json",
		"dictdata/jmdict_english/term_bank_21.json",
		"dictdata/jmdict_english/term_bank_22.json",
		"dictdata/jmdict_english/term_bank_23.json",
		"dictdata/jmdict_english/term_bank_24.json",
		"dictdata/jmdict_english/term_bank_25.json",
		"dictdata/jmdict_english/term_bank_26.json",
		"dictdata/jmdict_english/term_bank_27.json",
		"dictdata/jmdict_english/term_bank_28.json",
		"dictdata/jmdict_english/term_bank_29.json",
	}

	var slice Entries
	for _, dict := range dicts {
		jsonFile, err := os.Open(dict)
		if err != nil {
			fmt.Println(err)
		}
		byteValue, _ := ioutil.ReadAll(jsonFile)

		var entries [][]interface{}
		if err := json.Unmarshal(byteValue, &entries); err != nil {
			panic(err)
		}

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
			slice.data = append(slice.data, e)
		}
	}

	fmt.Println(slice.data[1315])

	// Server.JPI()
}
