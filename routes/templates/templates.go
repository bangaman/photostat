package templates

import (
    "html/template"
    "net/http"
    "os"
    "path/filepath"
)

func RenderTemplate(w http.ResponseWriter, tmpl []string, data interface{}) {
    cwd, _ := os.Getwd()
    t, err := template.ParseFiles(filepath.Join( cwd, "./routes/" + tmpl[0] + "/templates/" + tmpl[1] + ".html"))
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    err = t.Execute(w, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}