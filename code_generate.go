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

type MainData struct {
	Port string
}

type EnvData struct {
	Mongouri string
}

type setupData struct {
	DBname string
}

type InputData struct {
	Port string `json:"port"`
	Mongouri string `json:mongouri`
	DBname string `json:dbname`
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
	createMain(inputData.Port)
	createEnv(inputData.Mongouri)
	createRoute(nil)
	createResponse(nil)
	createModel(nil)
	createSetup(inputData.DBname)

	return c.JSON(http.StatusCreated,"Successfully code generated")
}

func createMain(Port string) {
	// Define the data for the template
	data := MainData{Port}

	// Generate the code file
	err := generateCodeFile("main.tmpl", "./output", "main.go", data)
	if err != nil {
		panic(err)
	}
}

func createEnv(Mongouri string) {
	// Define the data for the template
	data := EnvData{Mongouri}

	// Generate the code file
	err := generateCodeFile("env.tmpl", "./output", ".env", data)
	if err != nil {
		panic(err)
	}
}

func createRoute(data interface{}) {
	// Generate the code file
	err := generateCodeFile("route.tmpl", "./output/routes", "route.go", data)
	if err != nil {
		panic(err)
	}
}

func createResponse(data interface{}) {
	// Generate the code file
	err := generateCodeFile("response.tmpl", "./output/responses", "response.go", data)
	if err != nil {
		panic(err)
	}
}

func createModel(data interface{}) {
	// Generate the code file
	err := generateCodeFile("model.tmpl", "./output/models", "model.go", data)
	if err != nil {
		panic(err)
	}
}

func createSetup(DBname string) {
	// Define the data for the template
	data := EnvData{DBname}

	// Generate the code file
	err := generateCodeFile("setup.tmpl", "./output/configs", "setup.go", data)
	if err != nil {
		panic(err)
	}
}

func main() {
	e := echo.New()

	e.POST("/input", getInputs)

	e.Logger.Fatal(e.Start(":8080"))
}
