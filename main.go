package main

import (
	"fmt"
	"log"
	"net/http"
	"zealthy-helpdesk-backend/dao"
	"zealthy-helpdesk-backend/handler"
	"zealthy-helpdesk-backend/utility"
)

func main() {
	fmt.Println("Service started")

	config, err := utility.LoadApplicationConfig("config", "deploy.yml")
	if err != nil {
		log.Fatal(err)
	}

	dao.DbInit(&config.Postgres)
	log.Fatal(http.ListenAndServe(":8080", handler.InitRouter(&config.JWT)))
}
