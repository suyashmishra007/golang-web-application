package render

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

// TODO: Make template cache -> s3 [30-end]
func RenderTemplate(w http.ResponseWriter, tmpl string) error {
	templates, err := template.ParseFiles(
		filepath.Join("templates", "base.layout.tmpl"),
		filepath.Join("templates", tmpl),
	)
	if err != nil {
		fmt.Printf("Error parsing templates: %v\n", err)
		return err
	}
	err = templates.ExecuteTemplate(w, "base.layout.tmpl", nil)
	if err != nil {
		fmt.Printf("Error executing template: %v\n", err)
		return err
	}
	return nil
}
