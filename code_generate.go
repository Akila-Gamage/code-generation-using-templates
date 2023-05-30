package main

import (
	"net/http"
	"os"
	"path/filepath"
	"text/template"

	"github.com/labstack/echo/v4"
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

func getInputs(c echo.Context) error{
	createMain("6050")
	createEnv()
	createRoute()

	return c.JSON(http.StatusCreated,"Successfully code generated")
}

func createMain(Port string) {
	// Define the data for the template
	data := Data{Port}

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
	e := echo.New()
	e.POST("/input", getInputs)



	e.Logger.Fatal(e.Start(":8080"))
}
