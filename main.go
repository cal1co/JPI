package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/twuillemin/levenshteinsearch/pkg/levenshteinsearch"
	// Server "github.com/cal1co/jpi/server"
)

const (
	terms = "dictdata/jmdict_english/term_bank_1.json"
)

type Entry struct {
	Word      string
	Alternate string
	Form      string
	Blank     string
	Id        string
	Def       string
	SecondId  string
	Source    string
}
type Entries struct {
	Entries []Entry
}

func main() {

	dict := levenshteinsearch.CreateDictionary()

	jsonFile, err := os.Open(terms)
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var entries [][]interface{}
	if err := json.Unmarshal(byteValue, &entries); err != nil {
		panic(err)
	}

	// fmt.Println(entries[5010][5])
	// val := fmt.Sprint(entries[5010][5])
	// fmt.Println(val)
	// dict.Put(val)
	// dict.Put("interview")
	// fmt.Println(dict.WordCount)
	// dict.Put("interviewer")

	// dictMap := make(map[string]map)
	// for i := 0; i < len(entries); i++ {
	// 	dictMap[entries[i][0]] = {

	// 	}
	// }

	// for i := 0; i < len(entries); i++ {
	// 	dict.Put(fmt.Sprint("%i:", entries[i][5]))
	// }

	// var s []map[string]map[string]string
	// for i := 0; i < len(entries); i++ {
	// 	key := fmt.Sprint(entries[i][0])
	// 	def := fmt.Sprint(entries[i][5])
	// 	reading := fmt.Sprint(entries[i][1])

	// 	resMap := make(map[string]map[string]string)
	// 	resMap[key] = map[string]string{}
	// 	resMap[key]["Reading"] = reading
	// 	resMap[key]["Definition"] = def
	// 	s = append(s, resMap)
	// }
	// fmt.Println(s[9999])

	var slice []Entry
	for i := 0; i < len(entries); i++ {
		word := fmt.Sprint(entries[i][0])
		alternate := fmt.Sprint(entries[i][1])
		form := fmt.Sprint(entries[i][2])
		blank := fmt.Sprint(entries[i][3])
		id := fmt.Sprint(entries[i][4])
		def := fmt.Sprint(entries[i][5])
		secondid := fmt.Sprint(entries[i][6])
		source := fmt.Sprint(entries[i][7])
		e := Entry{
			Word:      word,
			Alternate: alternate,
			Form:      form,
			Blank:     blank,
			Id:        id,
			Def:       def,
			SecondId:  secondid,
			Source:    source,
		}
		slice = append(slice, e)
		dict.Put(def)
	}

	fmt.Println(slice[9876].Def)

	// jsonFile.Close()

	fmt.Println(dict.SearchAll("[cold]", 1))
	// Server.JPI()
}
