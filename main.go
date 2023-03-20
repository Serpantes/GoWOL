package main

import (
	"github.com/Serpantes/GoWOL/config"
	"github.com/Serpantes/GoWOL/server"
)
func main() {

	config.InitConfig()
	config.InitLogger()

	server.InitServer()
}
