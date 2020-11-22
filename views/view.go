package views

import (
	"text/template"
	"net/http"
	"path/filepath"
)

var (
	LayoutDir string = "views/layouts/"
	TemplateExt string = ".gohtml"
	TemplateDir string = "views/"
)


func NewView(layout string, files ...string) *View {
	addTemplatePath(files)
	addTemplateExt(files)
	files = append(files,layoutFiles()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
	Template: t,
	Layout: layout,
	}

}

//View new view
type View struct {
	Template *template.Template
	Layout string
}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := v.Render(w, nil); err != nil {
		panic(err)
	}
}

//Render is used torender the view with predefined template
func (v *View) Render(w http.ResponseWriter, data interface{}) error{
	w.Header().Set("Content-Type", "text/html")
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}


// Layout files return a slice of string representing
// the layout files used in our appliaction
func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err !=nil {
		panic(err)
	}
	return files
}

//Takes a slice of string representing file path for template
func addTemplatePath(files []string) {
	for i, f := range files {
		files[i] = TemplateDir + f
	}
}


func addTemplateExt(files []string) {
	for i, f := range files {
		files[i] = f + TemplateExt
	}
}
