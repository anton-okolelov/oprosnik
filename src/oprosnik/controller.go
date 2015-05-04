package oprosnik

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"oprosnik/model"
	"regexp"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	session := model.GetUserSession(w, r)
	if session.IsLogged {
		w.Write([]byte("welcome, " + session.Name))
	} else {
		render(w, "select-name.html")
	}
}

func admin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	render(w, "admin-form.html")
}

func adminSaveWords(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	text := r.FormValue("words")

	words := regexp.MustCompile("\r\n").Split(text, 1000)
	model.SaveSentences(words)
	http.Redirect(w, r, "/admin", 302)
}

func saveUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := r.FormValue("name")

	session := model.GetUserSession(w, r)
	session.Name = name
	session.Save()
	http.Redirect(w, r, "/", 302)
}
