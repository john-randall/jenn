package main

import (
	"embed"
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/iancoleman/strcase"
)

var (
	//go:embed templates/*.hbs
	tmpls embed.FS
)

func check(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func main() {

	args := os.Args[1:]
	command := args[0]

	path, err := os.Getwd()
	check(err)

	fmt.Println(path)

	switch command {
	case "component", "c":
		componentName := args[1]
		generateComponent(path, componentName)
	default:
		fmt.Println("Command not found")
	}
}

func generateComponent(path string, name string) {
	properName := strcase.ToCamel(name)
	baseOutputDir := path + "/" + properName

	// Make Directory for component
	os.MkdirAll(baseOutputDir, os.ModePerm)

	// Read Template Files and Create Templates
	t, err := template.ParseFS(tmpls, "templates/*")
	check(err)

	templateToFile(baseOutputDir+"/"+properName+".tsx", t, "component.hbs", properName)
	templateToFile(baseOutputDir+"/index.ts", t, "index.hbs", properName)

}

func templateToFile(outputFilePath string, template *template.Template, templateName string, data interface{}) {
	outputFile, err := os.Create(outputFilePath)
	check(err)

	defer outputFile.Close()

	err = template.ExecuteTemplate(outputFile, templateName, data)
	check(err)

	outputFile.Sync()
}
