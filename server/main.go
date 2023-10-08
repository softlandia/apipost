// Package main - 
package main

import (
	"context"
	"log"
	"main/app/control"
	"main/app/model"
	"net/http"

	"main/app/server"

	"github.com/julienschmidt/httprouter"
)

func main() {

	r := httprouter.New()
	routes(r)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}

	mdl := model.GetStorage()

	postgreSqlClient, err := server.NewClient(context.TODO(), 3, mdl.Storage)
	if err != nil {
		log.Fatal(err)
	}
}

func routes(r *httprouter.Router) {
	//путь к папке со внешними файлами: html, js, css, изображения и т.д.
	r.ServeFiles("/app/public/*filepath", http.Dir("public"))
	//что следует выполнять при входящих запросах указанного типа и по указанному адресу
	r.GET("/", control.StartPage)
	r.GET("/users", control.GetUsers)
}
