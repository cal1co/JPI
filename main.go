package main

import (
	"fmt"

	algo "github.com/cal1co/jpi/algo"
)

// Server "github.com/cal1co/jpi/server"

type Entry struct {
	Word      string
	Alternate string
	Id        int
	Def       []interface{}
	Dist      float32
}

type Entries struct {
	data []Entry
}

func main() {
	// dict := []string{
	// 	"dictdata/jmdict_english/term_bank_1.json",
	// 	"dictdata/jmdict_english/term_bank_2.json",
	// 	"dictdata/jmdict_english/term_bank_3.json",
	// 	"dictdata/jmdict_english/term_bank_4.json",
	// 	"dictdata/jmdict_english/term_bank_5.json",
	// 	"dictdata/jmdict_english/term_bank_6.json",
	// 	"dictdata/jmdict_english/term_bank_7.json",
	// 	"dictdata/jmdict_english/term_bank_8.json",
	// 	"dictdata/jmdict_english/term_bank_9.json",
	// 	"dictdata/jmdict_english/term_bank_10.json",
	// 	"dictdata/jmdict_english/term_bank_11.json",
	// 	"dictdata/jmdict_english/term_bank_12.json",
	// 	"dictdata/jmdict_english/term_bank_13.json",
	// 	"dictdata/jmdict_english/term_bank_14.json",
	// 	"dictdata/jmdict_english/term_bank_15.json",
	// 	"dictdata/jmdict_english/term_bank_16.json",
	// 	"dictdata/jmdict_english/term_bank_17.json",
	// 	"dictdata/jmdict_english/term_bank_18.json",
	// 	"dictdata/jmdict_english/term_bank_19.json",
	// 	"dictdata/jmdict_english/term_bank_20.json",
	// 	"dictdata/jmdict_english/term_bank_21.json",
	// 	"dictdata/jmdict_english/term_bank_22.json",
	// 	"dictdata/jmdict_english/term_bank_23.json",
	// 	"dictdata/jmdict_english/term_bank_24.json",
	// 	"dictdata/jmdict_english/term_bank_25.json",
	// 	"dictdata/jmdict_english/term_bank_26.json",
	// 	"dictdata/jmdict_english/term_bank_27.json",
	// 	"dictdata/jmdict_english/term_bank_28.json",
	// 	"dictdata/jmdict_english/term_bank_29.json",
	// }

	// var slice Entries
	// for _, dictBank := range dict {
	// 	jsonFile, err := os.Open(dictBank)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	byteValue, _ := ioutil.ReadAll(jsonFile)

	// 	var entries [][]interface{}
	// 	if err := json.Unmarshal(byteValue, &entries); err != nil {
	// 		panic(err)
	// 	}

	// 	for i := 0; i < len(entries); i++ {
	// 		word := fmt.Sprint(entries[i][0])
	// 		alternate := fmt.Sprint(entries[i][1])
	// 		id := fmt.Sprint(entries[i][4])
	// 		intId, err := strconv.Atoi(id)
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		def, ok := entries[i][5].([]interface{})
	// 		if !ok {
	// 			panic(ok)
	// 		}
	// 		e := Entry{
	// 			Word:      word,
	// 			Alternate: alternate,
	// 			Id:        intId,
	// 			Def:       def,
	// 		}
	// 		slice.data = append(slice.data, e)
	// 	}
	// }
	// Server.JPI()

	fmt.Println(algo.Blah())
}
