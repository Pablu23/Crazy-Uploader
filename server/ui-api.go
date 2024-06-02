package server

import (
	"net/http"
	"text/template"

	"github.com/pablu23/uploader/view"
	"github.com/pablu23/uploader/view/viewmodels"
)

func (s *Server) GetTest(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(view.GetComponent(view.ListItem))
	viewmodel := viewmodels.NewListViewmodel(viewmodels.ListItem{Name: "Zam", Age: 20}, viewmodels.ListItem{Name: "Lukas", Age: 20})
	tmpl.Execute(w, viewmodel)
}

func (s *Server) GetList(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(view.GetViewTemplate(view.List))
	tmpl.Execute(w, nil)
}
