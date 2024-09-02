package main

import (
	"to_do_list_app/app/api"
	"to_do_list_app/config"
	"to_do_list_app/db"
)

func init() {
	config.InitEnvs()
	db.ConnectToDB()
}

func main() {
	api.RunServer()
}
