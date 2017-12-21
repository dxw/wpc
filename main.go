package main

import (
	"bufio"
	"bytes"
	"fmt"
	flag "github.com/spf13/pflag"
	"io/ioutil"
	"os"
)

//Filemodes
const (
	NOEXEC = 0644
	EXEC   = 0755
)

func main() {
	multisite := flag.Bool("multisite", false, "Make it a multisite")
	flag.Parse()
	project := flag.Arg(0)
	if project == "" {
		fmt.Println("You must set a value for the project argument")
		os.Exit(1)
	}

	dirs := []string{"bin", "setup/content"}
	makeDirs(dirs)

	createDockerCompose(project)
	creating("bin/wp", []byte(WPCONTENT), EXEC)
	creating("bin/console", []byte(CONSOLECONTENT), EXEC)
	creating("bin/setup", []byte(SETUPCONTENT), EXEC)
	creating("setup/external.sh", []byte(EXTERNALCONTENT), EXEC)
	createInternal(*multisite)
}

func createDockerCompose(project string) {
	dockerComposeContents := []byte(DOCKERCOMPOSECONTENT)
	dockerComposeContents = findAndReplace(dockerComposeContents, []byte("!!!PROJECTNAME!!!"), []byte(project))
	creating("docker-compose.yml", dockerComposeContents, NOEXEC)
}

func createInternal(multisite bool) {
	contents := []byte(INTERNALCONTENT)
	if multisite {
		contents = findAndReplace(contents, []byte("!!!INSTALLTYPE!!!"), []byte("multisite-install"))
		contents = findAndReplace(contents, []byte("!!!ACTIVATIONTYPE!!!"), []byte("--network"))
		contents = findAndReplace(contents, []byte("!!!THEMEENABLE!!!"), []byte("wp theme enable --network $theme"))
	} else {
		contents = findAndReplace(contents, []byte("!!!INSTALLTYPE!!!"), []byte("install"))
		contents = findAndReplace(contents, []byte("!!!ACTIVATIONTYPE!!!"), []byte(""))
		contents = findAndReplace(contents, []byte("!!!THEMEENABLE!!!"), []byte(""))
	}
	creating("setup/internal.sh", contents, EXEC)
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

func findAndReplace(haystack []byte, old []byte, new []byte) []byte {
	newHaystack := bytes.Replace(haystack, old, new, -1)
	return newHaystack
}
