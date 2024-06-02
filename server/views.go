package server

import (
	"net/http"
	"text/template"

	"github.com/pablu23/uploader/view"
)

func (s *Server) GetList(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(view.GetViewTemplate(view.List))
	tmpl.Execute(w, nil)
}
