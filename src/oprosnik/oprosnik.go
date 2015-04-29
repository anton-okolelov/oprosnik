package oprosnik

import( 
    "net/http"
	"github.com/julienschmidt/httprouter"
	"log"
)

func Start() {
	
    router := httprouter.New()
    router.GET("/", index)
    router.GET("/admin/", admin)
	router.POST("/admin/save", adminSaveWords)
	router.POST("/save-name", saveUserName)

    log.Fatal(http.ListenAndServe(":8080", router))
}