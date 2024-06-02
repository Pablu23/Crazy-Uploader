package server

import (
	"net/http"
	"text/template"
	"time"

	"github.com/pablu23/uploader/view"
	"github.com/pablu23/uploader/view/components"
)

func (s *Server) GetTest(w http.ResponseWriter, r *http.Request) {
  time.Sleep(5 * time.Second)
	tmpl := template.Must(view.GetComponent(view.ListItem))
	viewmodel := components.NewListComponent(components.ListItem{Name: "Zam", Age: 20}, components.ListItem{Name: "Lukas", Age: 20})
	tmpl.Execute(w, viewmodel)
}
