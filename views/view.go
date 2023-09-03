package views

import (
	"html/template"
	"net/http"
)

type View struct {
	template *template.Template
}

func NewView() *View {
	return &View{
		template: template.Must(template.ParseGlob("./**/*.html")),
	}
}

func (v *View) Write(w http.ResponseWriter, view string, data any) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	v.template.ExecuteTemplate(w, view+".html", data)
}
