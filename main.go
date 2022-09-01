package main

import (
	data "github.com/cal1co/jpi/handleData"
	Server "github.com/cal1co/jpi/server"
)

func main() {
	res := data.InitStructures()
	Server.JPI(res)
}
