package main

import (
	"oprosnik"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"	
)


func main() {
	log.Println("Starting...")

	router := httprouter.New()

	router.GET("/", oprosnik.Index)
	router.GET("/admin/", oprosnik.Admin)
	router.POST("/admin/save", oprosnik.AdminSaveSentences)
	router.POST("/save-name", oprosnik.SaveUserName)
	router.POST("/answer", oprosnik.SaveAnswer)
	
	// статика
	router.GET("/css/*filepath", oprosnik.StaticFiles)
	router.GET("/js/*filepath", oprosnik.StaticFiles)
	
	// обрабатываем фатальные ошибки
	router.PanicHandler = oprosnik.PanicHandler

	// запускаем сервер
	log.Fatal(http.ListenAndServe(":8080", context.ClearHandler(router)))
}
