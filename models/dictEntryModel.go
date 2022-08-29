package main

type entries struct {
	entres []entry
}

type entry struct {
	Word      string
	Alternate string
	Form      string
	Blank     string
	Id        int
	Def       string
}
