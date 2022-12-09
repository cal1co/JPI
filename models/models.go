package models

type Entries struct {
	entries []Entry
}

type Entry struct {
	Word      string
	Alternate string
	Form      string
	Blank     string
	Id        int
	Def       string
}
