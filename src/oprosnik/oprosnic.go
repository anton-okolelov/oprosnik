package oprosnik

import(
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "log"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "Опросник!\n")
}

func admin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	render(w, "admin-form.html")
}

func save(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func Start() {
	
    router := httprouter.New()
    router.GET("/", index)
    router.GET("/admin/", admin)
	router.POST("/admin/save", save)

    log.Fatal(http.ListenAndServe(":8080", router))
}