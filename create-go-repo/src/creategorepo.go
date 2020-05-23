package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// GlobalVars is..
type GlobalVars struct {
	lang            string
	isDockerEnabled bool
	folderPath      string
	repoName        string
	errorMsg        string
}

var _globalVars GlobalVars

func printGlobalVars() {
	fmt.Println(" Language:" + _globalVars.lang)
	fmt.Println(" Repo Name:" + _globalVars.repoName)
	fmt.Println(" Folder Path:" + _globalVars.folderPath)
	fmt.Println(" Docker Enabled:" + strconv.FormatBool(_globalVars.isDockerEnabled))
}

func main() {
	initialize()
	createFolder()
}

func createFolder() {
	dirPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		logFatalError(err)
	}

	if len(_globalVars.folderPath) != 0 {
		dirPath = _globalVars.folderPath + "\\" + _globalVars.repoName
		_, err := os.Stat(dirPath)
		if os.IsNotExist(err) == false {
			_globalVars.errorMsg = ">>Error: Folder already exists."
			exitProg()
		}
	} else {
		dirPath = dirPath + "\\" + _globalVars.repoName
	}

	fmt.Println("-- Creating the folder:" + dirPath)
	err = os.MkdirAll(dirPath, 0777)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("-- Folder created successfully --")
}

func initialize() {
	println(" -- INITIALIZING --")
	_globalVars.lang = ""
	_globalVars.isDockerEnabled = false
	_globalVars.folderPath = ""
	_globalVars.errorMsg = ""

	args := os.Args[0:]
	for i := 1; i < len(args); i++ {
		command := args[i]
		if command == "-d" {
			_globalVars.isDockerEnabled = true
		}
		if strings.HasPrefix(command, "-lang") {
			_globalVars.lang = getValue(command, "-lang")
		}
		if strings.HasPrefix(command, "-name") {
			_globalVars.repoName = getValue(command, "-name")
		}
		if strings.HasPrefix(command, "-fp") {
			_globalVars.folderPath = getValue(command, "-fp")
		}
		if command == "-help" {
			showHelp()
			exitProg()
		}
	}
	validateInputParams()
}

func validateInputParams() {
	fmt.Println("-- Validating Inputs --")
	printGlobalVars()
	if len(_globalVars.lang) == 0 ||
		len(_globalVars.repoName) == 0 {
		_globalVars.errorMsg = ">> Error: Invalid params. Lang / name missing"
		exitProg()
	}

	fmt.Println("-- Validation Complete. No Errors --")
}

func getValue(cmdString string, cmdPrefix string) string {
	cmd := strings.Replace(cmdString, cmdPrefix+":", "", 1)
	l := len(cmd)
	if l == 0 {
		_globalVars.errorMsg = ">> Error: Invalid " + cmdPrefix
		log.Fatal(_globalVars.errorMsg)
		exitProg()
		return ""
	}

	return cmd
}

func logFatalError(err error) {
	log.Fatal(err)
	os.Exit(3)
}
func exitProg() {
	fmt.Println(_globalVars.errorMsg)
	os.Exit(3)
}

func createGoRepo() {

}

func showHelp() {
	fmt.Printf(`
	-name 	>>	{Required}
				Example_>> -name:<repoName>
				Indicates Name of the Project.
				This is the name that is used as a repo in github.

	-lang	>> 	{Required}
				Example_>> -lang:<language>
				example:
					createrepo -lang:go
					createrepo -lang:python
					etc..

	-d		>> 	{Optional}
				Indicates if repo needs to be dockerized.
				Adds a Dockerfile.

	-fp		>> 	{Optional}
				Example_>> -fp:<folderpath>; 
				Indicates parent folder path.
				example:
					createrepo -fp:c:\github\projects
	
	-help	>>	{Optional}
				Shows Help.
	`)
}
