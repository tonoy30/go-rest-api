package utils

import (
	"fmt"
	"net/http"
	"text/template"
)

func ServeTemplate(data interface{}, fileName string, w http.ResponseWriter) {
	file := fmt.Sprintf("./templates/%s.html", fileName)
	tmpl := template.Must(template.ParseFiles(file))
	tmpl.Execute(w, data)
}
