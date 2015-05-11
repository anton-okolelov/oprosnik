package main

import (
	"oprosnik"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"oprosnik/model"
)


func main() {
	log.Println("Starting...")

	// в отдельном потоке периодически чистим старые сессии 
	go model.SessionGarbageCollector()

	router := httprouter.New()

	router.GET("/", oprosnik.Index)
	router.GET("/question", oprosnik.Question)
	router.GET("/admin/", oprosnik.Admin)
	router.POST("/admin/save", oprosnik.AdminSaveSentences)
	router.POST("/admin/cleanresults", oprosnik.AdminCleanResults)
	router.POST("/save-name", oprosnik.SaveUserName)
	router.POST("/answer", oprosnik.SaveAnswer)
	
	// статика
	router.GET("/results/*filepath", oprosnik.StaticFiles)
	router.GET("/bower_components/*filepath", oprosnik.StaticFiles)
	
	// обрабатываем фатальные ошибки, error 500
	router.PanicHandler = oprosnik.PanicHandler

	// запускаем сервер
	log.Fatal(http.ListenAndServe(":8080", context.ClearHandler(router)))
}
