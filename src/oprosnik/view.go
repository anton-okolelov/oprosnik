package oprosnik

import (
	"net/http"
	"text/template"
)

var templatesPath = "resources/templates"
var templates = make(map[string]*template.Template)

// компилируем шаблоны, причем каждый в паре с base.html,
// чтобы организовать как бы наследование. Вообще, надо какую-то библиотеку найти
func init() {
	templateNames := []string{"admin-form.html", "select-name.html", "question.html", "okay.html"}
	for _, templateName := range templateNames {
		path := templatesPath + "/" + templateName
		base := templatesPath + "/base.html"
		templates[templateName] = template.Must(template.ParseFiles(path, base))
	}
}

// рендерим относледованный шаблон
func renderExtended(w http.ResponseWriter, name string, data interface{}) {
	templates[name].ExecuteTemplate(w, "base", data)
}
