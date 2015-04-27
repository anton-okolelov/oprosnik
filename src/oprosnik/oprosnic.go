package oprosnik

import(
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "log"
	//"strings"
	"encoding/json"
	"os"
	"path/filepath"
	"regexp"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "Опросник!\n")
}

func admin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	render(w, "admin-form.html")
	
}

func save(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    text := r.FormValue("words")
	
	words := regexp.MustCompile("\r\n").Split(text, 1000)
	wordsJson, _ := json.Marshal(words)
	f, err := os.Create("data/questions.json")
	if err != nil {
		path, _ := filepath.Abs(".")
		log.Fatal("File  " + path, err)
	}
	f.Write(wordsJson)
	f.Close()
	http.Redirect(w, r, "/admin", 302)
	
}

func Start() {
	
    router := httprouter.New()
    router.GET("/", index)
    router.GET("/admin/", admin)
	router.POST("/admin/save", save)

    log.Fatal(http.ListenAndServe(":8080", router))
}