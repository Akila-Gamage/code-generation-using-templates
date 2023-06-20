package main

import (
	"net/http"
	"os"
	"path/filepath"
	"text/template"
	"encoding/json"
	"io/ioutil"

	"github.com/labstack/echo/v4"
)

//struct to represent the input data
type InputData struct {
	Port 			string `json:"port"`
	MongoUri 		string `json:"mongouri"`
	DBname 			string `json:"dbname"`
	CollectionName  string `json:"collectionname"`
}

func generateCodeFile(templatePath string, folderPath string, outputPath string, data interface{}) error {
	// Load the template file
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	// Create the output file 
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

	// Execute the template 
	err = tmpl.Execute(outputFile, data)
	if err != nil {
		return err
	}

	return nil
}

func getInputs(c echo.Context) error{
	// Read the request body
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}

	// Parse the JSON data into the InputData struct
	var inputData InputData
	err = json.Unmarshal(body, &inputData)
	if err != nil {
		return err
	}

	// Use the inputData in your code generation functions
	createMain(inputData)
	createEnv(inputData)
	createRoute(nil)
	createResponse(nil)
	createModel(nil)
	createSetup(inputData)
	createEnvc(nil)
	createController(inputData)

	return c.JSON(http.StatusCreated,"Successfully code generated")
}

func createMain(data InputData) {
	// Generate the code file
	err := generateCodeFile("./templates/main.tmpl", "/Users/akilagamage/Documents/go-workspace/CrudWithGoGenerate/crud-with-go", "main.go", data)
	if err != nil {
		panic(err)
	}
}

func createEnv(data InputData) {
	// Generate the code file
	err := generateCodeFile("./templates/env.tmpl", "/Users/akilagamage/Documents/go-workspace/CrudWithGoGenerate/crud-with-go", ".env", data)
	if err != nil {
		panic(err)
	}
}

func createRoute(data interface{}) {
	// Generate the code file
	err := generateCodeFile("./templates/route.tmpl", "/Users/akilagamage/Documents/go-workspace/CrudWithGoGenerate/crud-with-go/routes", "route.go", data)
	if err != nil {
		panic(err)
	}
}

func createResponse(data interface{}) {
	// Generate the code file
	err := generateCodeFile("./templates/response.tmpl", "/Users/akilagamage/Documents/go-workspace/CrudWithGoGenerate/crud-with-go/responses", "response.go", data)
	if err != nil {
		panic(err)
	}
}

func createModel(data interface{}) {
	// Generate the code file
	err := generateCodeFile("./templates/model.tmpl", "/Users/akilagamage/Documents/go-workspace/CrudWithGoGenerate/crud-with-go/models", "model.go", data)
	if err != nil {
		panic(err)
	}
}

func createSetup(data InputData) {
	// Generate the code file
	err := generateCodeFile("./templates/setup.tmpl", "/Users/akilagamage/Documents/go-workspace/CrudWithGoGenerate/crud-with-go/configs", "setup.go", data)
	if err != nil {
		panic(err)
	}
}

func createEnvc(data interface{}) {
	// Generate the code file
	err := generateCodeFile("./templates/envc.tmpl", "/Users/akilagamage/Documents/go-workspace/CrudWithGoGenerate/crud-with-go/configs", "env.go", data)
	if err != nil {
		panic(err)
	}
}

func createController(data InputData) {
	// Generate the code file
	err := generateCodeFile("./templates/controller.tmpl", "/Users/akilagamage/Documents/go-workspace/CrudWithGoGenerate/crud-with-go/controllers", "controller.go", data)
	if err != nil {
		panic(err)
	}
}

func main() {
	e := echo.New()

	e.POST("/input", getInputs)

	e.Logger.Fatal(e.Start(":6050"))
}
