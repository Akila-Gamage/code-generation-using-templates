package main

import (
	"os"
	"text/template"
	"path/filepath"
)

type Data struct {
	Port string
}

func generateCodeFile(templatePath string, folderPath string, outputPath string, data interface{}) error {
	// Load the template file
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	// Create the output file in the desired folder
	outputFolderPath := folderPath // Specify the desired folder path here
	err = os.MkdirAll(outputFolderPath, os.ModePerm)
	if err != nil {
		panic(err)
	}
	outputFilePath := filepath.Join(outputFolderPath, outputPath)
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()


	// Execute the template with the provided data and write the output to the file
	err = tmpl.Execute(outputFile, data)
	if err != nil {
		return err
	}

	return nil
}

func createMain() {
	// Define the data for the template
	data := Data{Port: "6060"}

	// Generate the code file
	err := generateCodeFile("main.tmpl", "./output", "main.go", data)
	if err != nil {
		panic(err)
	}
}

func createEnv() {
	// Define the data for the template
	data := Data{}

	// Generate the code file
	err := generateCodeFile("env.tmpl", "./output", ".env", data)
	if err != nil {
		panic(err)
	}
}

func createRoute() {
	// Define the data for the template
	data := Data{}

	// Generate the code file
	err := generateCodeFile("route.tmpl", "./output/routes", "route.go", data)
	if err != nil {
		panic(err)
	}
}

func createResponse() {
	// Define the data for the template
	data := Data{}

	// Generate the code file
	err := generateCodeFile("response.tmpl", "./output/responses", "response.go", data)
	if err != nil {
		panic(err)
	}
}

func main() {
	createMain()
	createEnv()
	createRoute()
}
