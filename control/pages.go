package control

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
)

func StartPage(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	path := filepath.Join("app", "public", "html", "startStaticPage.html")

	tmpl, err := template.ParseFiles(path)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
}
