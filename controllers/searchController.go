package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/cal1co/jpi/algo"
)

type Entry struct {
	Word      string
	Alternate string
	Freq      int
	Def       []interface{}
	Dist      float32
}

type Entries struct {
	data []Entry
}

type Output struct {
	Nodes algo.Trie
	Slice Entries
}

func Fetch(trie *algo.Trie, slice Entries, word string) []Entry {
	res := algo.Find(word, trie, 2, 0, 0)
	sort.Slice(res, func(i, j int) bool {
		return res[i].Score > res[j].Score
	})
	var returnData []Entry
	for i := 0; i < len(res); i++ {
		sort.Slice(res[i].Position, func(k, l int) bool {
			return slice.data[res[i].Position[k]].Freq > slice.data[res[i].Position[l]].Freq
		})
		for j := 0; j < len(res[i].Position); j++ {
			returnData = append(returnData, slice.data[res[i].Position[j]])
		}
	}
	return returnData
}

func InitStructures() Output {
	dict := []string{
		"dictdata/jmdict_english/term_bank_1.json",
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

	trie := algo.InitTrie()

	var slice Entries
	for _, dictBank := range dict {
		jsonFile, err := os.Open(dictBank)
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
			id := fmt.Sprint(entries[i][4])
			intId, err := strconv.Atoi(id)
			if err != nil {
				panic(err)
			}
			def, ok := entries[i][5].([]interface{})
			if !ok {
				panic(ok)
			}
			e := Entry{
				Word:      word,
				Alternate: alternate,
				Freq:      intId,
				Def:       def,
			}
			for j := 0; j < len(e.Def); j++ {
				definitions := strings.Split(strings.ToLower(fmt.Sprint(e.Def[j])), " ")
				algo.Insert(strings.ToLower(fmt.Sprint(e.Def[j])), len(slice.data), trie)
				if len(definitions) > 1 {
					for k := 0; k < len(definitions); k++ {
						algo.Insert(definitions[k], len(slice.data), trie)
					}
				}
			}
			slice.data = append(slice.data, e)
		}
	}
	var res Output
	res.Nodes = *trie
	res.Slice = slice
	return res
}
