package main

import (
	"bufio"
	"bytes"
	"fmt"
	flag "github.com/spf13/pflag"
	"io/ioutil"
	"os"
	"text/template"
)

//Filemodes
const (
	NOEXEC = 0644
	EXEC   = 0755
)

type Project struct {
	Name           string
	Multisite      bool
	InstallType    string
	ActivationType string
	ThemeEnable    string
}

func newProject(name string, multisite bool) Project {
	var installType string
	var activationType string
	var themeEnable string
	if multisite {
		installType = "multisite-install"
		activationType = "--network"
		themeEnable = "wp theme enable --network $theme"
	} else {
		installType = "install"
		activationType = ""
		themeEnable = ""
	}
	return Project{name, multisite, installType, activationType, themeEnable}
}

func main() {
	multisite := flag.Bool("multisite", false, "Make it a multisite")
	flag.Parse()
	projectName := flag.Arg(0)
	if projectName == "" {
		fmt.Println("You must set a value for the project argument")
		os.Exit(1)
	}

	project := newProject(projectName, *multisite)

	dirs := []string{"bin", "setup/content"}
	makeDirs(dirs)

	createFromTemplate(DOCKERCOMPOSECONTENT, project, "docker-compose.yml")
	creating("bin/wp", []byte(WPCONTENT), EXEC)
	creating("bin/console", []byte(CONSOLECONTENT), EXEC)
	creating("bin/setup", []byte(SETUPCONTENT), EXEC)
	creating("setup/external.sh", []byte(EXTERNALCONTENT), EXEC)
	createFromTemplate(INTERNALCONTENT, project, "setup/internal.sh")
}

func createFromTemplate(templateContent string, project Project, file string) {
	t := template.Must(template.New("content").Parse(templateContent))
	var contents []byte
	var tOutput bytes.Buffer
	err := t.Execute(&tOutput, project)
	if err == nil {
		contents = tOutput.Bytes()
		creating(file, contents, EXEC)
	}
}

func makeDirs(directories []string) {
	for _, directory := range directories {
		fmt.Println("-> Creating", directory)
		os.MkdirAll(directory, EXEC)
	}
}

func creating(file string, data []byte, fileMode os.FileMode) {
	if okToWrite(file) {
		fmt.Println("-> Creating", file)
		ioutil.WriteFile(file, data, fileMode)
	}
}

func okToWrite(file string) bool {
	if _, err := os.Stat(file); !os.IsNotExist(err) {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println(file, "already exists, overwrite (y/n)?")
		response, _ := reader.ReadString('\n')
		if response != "y\n" && response != "Y\n" {
			return false
		}
	}
	return true
}
