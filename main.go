//go:generate go run generate.go
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"

	flag "github.com/spf13/pflag"
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

	dirs := []string{"bin", "setup/content", "script"}
	makeDirs(dirs)

	createFromTemplate(rawTemplates["templates/docker-compose.yml.tmpl"], project, "docker-compose.yml", NOEXEC)
	creating("bin/wp", rawTemplates["templates/bin/wp.tmpl"], EXEC)
	creating("bin/console", rawTemplates["templates/bin/console.tmpl"], EXEC)
	creating("bin/setup", rawTemplates["templates/bin/setup.tmpl"], EXEC)
	creating("setup/external.sh", rawTemplates["templates/setup/external.sh.tmpl"], EXEC)
	createFromTemplate(rawTemplates["templates/setup/internal.sh.tmpl"], project, "setup/internal.sh", EXEC)

	// dxwRFC compliance https://github.com/dxw/tech-team-rfcs/pull/23
	creating("script/bootstrap", rawTemplates["templates/script/bootstrap.tmpl"], EXEC)
	creating("script/setup", rawTemplates["templates/script/setup.tmpl"], EXEC)
	creating("script/update", rawTemplates["templates/script/update.tmpl"], EXEC)
	creating("script/server", rawTemplates["templates/script/server.tmpl"], EXEC)
	creating("script/console", rawTemplates["templates/script/console.tmpl"], EXEC)

	if *multisite {
		makeDirs([]string{"config"})
		creating("config/server.php", rawTemplates["templates/config/server.php.tmpl"], NOEXEC)
	}
}

func createFromTemplate(templateContent string, project Project, file string, fileMode os.FileMode) {
	t := template.Must(template.New("content").Parse(templateContent))
	var tOutput bytes.Buffer
	err := t.Execute(&tOutput, project)
	if err == nil {
		creating(file, tOutput.String(), fileMode)
	}
}

func makeDirs(directories []string) {
	for _, directory := range directories {
		fmt.Println("-> Creating", directory)
		os.MkdirAll(directory, EXEC)
	}
}

func creating(file string, data string, fileMode os.FileMode) {
	if okToWrite(file) {
		fmt.Println("-> Creating", file)
		ioutil.WriteFile(file, []byte(data), fileMode)
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
