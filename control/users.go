package control

import (
	"html/template"
	"main/app/model"
	"net/http"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
)

func GetUsers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	users, err := model.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	userPath := filepath.Join("app", "public", "html", "userDynamicPage.html")
	commonPath := filepath.Join("app", "public", "html", "common.html")

	tmpl, err := template.ParseFiles(userPath, commonPath)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = tmpl.ExecuteTemplate(w, "users", users)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}
