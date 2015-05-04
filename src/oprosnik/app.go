package oprosnik

import (
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func Start() {

	router := httprouter.New()
	router.GET("/", index)
	router.GET("/admin/", admin)
	router.POST("/admin/save", adminSaveWords)
	router.POST("/save-name", saveUserName)
	router.POST("/save-answer", saveAnswer)

	log.Fatal(http.ListenAndServe(":8080", context.ClearHandler(router)))
}
