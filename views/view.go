package views

import (
	"html/template"
	"path/filepath"
)

var (
	LayoutDir string = "views/layouts/"
	TemplateExt string = ".gohtml"
)


func NewView(layout string, files ...string) *View {
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


// Layout files return a slice of string representing
// the layout files used in our appliaction
func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err !=nil {
		panic(err)
	}
	return files
}
