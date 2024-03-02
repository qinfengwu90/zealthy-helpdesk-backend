package main

import (
	"fmt"
	"log"
	"net/http"
	"zealthy-helpdesk-backend/dao"
	"zealthy-helpdesk-backend/handler"
)

func main() {
	fmt.Println("Service started")
	dao.DbInit()
	log.Fatal(http.ListenAndServe(":8080", handler.InitRouter()))
}
