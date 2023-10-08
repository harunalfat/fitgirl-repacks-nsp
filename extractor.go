package main

import (
	_ "embed"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

//go:embed nspBuild.exe
var binary []byte

const nspBuildExecutable = "./nspBuild.exe"

func extractNcaFilesPath(directoryPath string) []string {
	filesPath := make([]string, 0)
	entries, err := os.ReadDir(directoryPath)
	if err != nil {
		panic(err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			splitted := strings.Split(entry.Name(), ".")
			if len(splitted) > 1 && strings.ToLower(splitted[len(splitted)-1]) == "nca" {
				fileName := directoryPath + "/" + entry.Name()
				fmt.Printf("Fetching file %s\n", entry.Name())
				filesPath = append(filesPath, fileName)
			}

			continue
		}

		dirName := directoryPath + "/" + entry.Name()
		fmt.Printf("Getting into directory %s\n", dirName)
		recResults := extractNcaFilesPath(dirName)
		filesPath = append(filesPath, recResults...)
	}

	return filesPath
}

func writeNspBuildExecutable() {
	err := os.WriteFile(nspBuildExecutable, binary, 0755)
	if err != nil {
		panic(err)
	}
}

func main() {
	outputName := flag.String("o", "out.nsp", "NSP output file name. Example: super_smash_bros.nsp")
	flag.Parse()

	execPath, err := os.Executable()
	if err != nil {
		panic(err)
	}

	pwd := filepath.Dir(execPath)

	nandRegisteredContentsPath := pwd + "/user/nand/user/Contents/registered"
	// List of current directory children (files or folders)
	_, err = os.ReadDir(nandRegisteredContentsPath)
	if err != nil {
		panic(errors.New("NAND NCA contents cannot be found. Wrong location or not compatible with Fitgirl-Repacks switch emulated format"))
	}

	arguments := extractNcaFilesPath(nandRegisteredContentsPath)

	writeNspBuildExecutable()
	output, err := exec.Command(nspBuildExecutable, append([]string{*outputName}, arguments...)...).Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))
	fmt.Printf("Succeed build file %s\n", *outputName)

}
