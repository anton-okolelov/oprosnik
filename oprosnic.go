package main

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "log"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "Опросник!\n")
}

func Admin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	tmpl, _ := template.ParseFiles("templates/admin-form.html")
	_ = tmpl.Execute(w, nil)
}

func Save(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func main() {
    router := httprouter.New()
    router.GET("/", Index)
    router.GET("/admin/", Admin)
	router.POST("/admin/save", Save)

    log.Fatal(http.ListenAndServe(":8080", router))
}