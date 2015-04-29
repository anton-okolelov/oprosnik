package oprosnik

import(
    "net/http"
	"regexp"
	"oprosnik/model"
	"github.com/julienschmidt/httprouter"	
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    	render(w, "select-name.html")
}

func admin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	render(w, "admin-form.html")
}

func adminSaveWords(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    text := r.FormValue("words")
	
	words := regexp.MustCompile("\r\n").Split(text, 1000)	
	model.SaveWords(words)
	http.Redirect(w, r, "/admin", 302)	
}


func saveUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    text := r.FormValue("words")
	
	words := regexp.MustCompile("\r\n").Split(text, 1000)	
	model.SaveWords(words)
	http.Redirect(w, r, "/admin", 302)	
}