package view

import "text/template"

func GetComponent(component Component) (*template.Template, error) {
	var path string
	switch component {
	case ListItem:
		path = "view/components/list.html"
	}
	return template.ParseFiles(path)
}

func GetViewTemplate(view View) (*template.Template, error) {
	var path string
	switch view {
	case List:
		path = "view/views/list.html"
	}
	return template.ParseFiles(path)
}
