package main

import (
	"bufio"
	"bytes"
	"fmt"
	flag "github.com/spf13/pflag"
	"io/ioutil"
	"os"
)

func main() {
	multisite := flag.Bool("multisite", false, "Make it a multisite")
	flag.Parse()
	project := flag.Arg(0)
	if project == "" {
		fmt.Println("You must set a value for the project argument")
		os.Exit(1)
	}

	makedir("bin")
	makedir("setup")
	makedir("setup/content")

	createDockerCompose(project)
	createBinWp()
	createBinConsole()
	createBinSetup()
	createExternal()
	createInternal(*multisite)
}

func createDockerCompose(project string) {
	dockerComposeContents := dockerComposeContents()
	dockerComposeContents = findAndReplace(dockerComposeContents, []byte("!!!PROJECTNAME!!!"), []byte(project))
	creatingNoexec("docker-compose.yml", dockerComposeContents)
}

func createBinWp() {
	contents := wpContents()
	creating("bin/wp", contents)
}

func createBinConsole() {
	contents := consoleContents()
	creating("bin/console", contents)
}

func createBinSetup() {
	contents := setupContents()
	creating("bin/setup", contents)
}

func createExternal() {
	contents := externalContents()
	creating("setup/external.sh", contents)
}

func createInternal(multisite bool) {
	contents := internalContents()
	if multisite {
		contents = findAndReplace(contents, []byte("!!!INSTALLTYPE!!!"), []byte("multisite-install"))
		contents = findAndReplace(contents, []byte("!!!ACTIVATIONTYPE!!!"), []byte("--network"))
		contents = findAndReplace(contents, []byte("!!!THEMEENABLE!!!"), []byte("wp theme enable --network $theme"))
	} else {
		contents = findAndReplace(contents, []byte("!!!INSTALLTYPE!!!"), []byte("install"))
		contents = findAndReplace(contents, []byte("!!!ACTIVATIONTYPE!!!"), []byte(""))
		contents = findAndReplace(contents, []byte("!!!THEMEENABLE!!!"), []byte(""))
	}
	creating("setup/internal.sh", contents)
}

func makedir(directory string) {
	fmt.Println("-> Creating", directory)
	os.Mkdir(directory, 0755)
}

func creatingNoexec(file string, data []byte) {
	if okToWrite(file) {
		fmt.Println("-> Creating", file)
		ioutil.WriteFile(file, data, 0644)
	}
}

func creating(file string, data []byte) {
	creatingNoexec(file, data)
	os.Chmod(file, 0755)
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
