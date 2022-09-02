package main

import (
	controller "github.com/cal1co/jpi/controllers"
	Server "github.com/cal1co/jpi/server"
)

func main() {
	res := controller.InitStructures()
	Server.JPI(res)
}
