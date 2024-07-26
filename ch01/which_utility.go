package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// First Part: Read commandline argument
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Println("Not enough arguments")
		return
	}

	// Second Part: Read PATH variable, splitting it and iterating over dir in PATH
	filename := arguments[1]
	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)

	// Third path: loop through all file in each dir, check that if it's exe file and match with name
	for _, dir := range pathSplit {
		fullPath := filepath.Join(dir, filename)
		// Check if exist or not?
		fileInfor, err := os.Stat(fullPath)
		if err != nil {
			continue
		}
		mode := fileInfor.Mode()
		// Is a regular file?
		if !mode.IsRegular() {
			continue
		}
		// Is exe file?
		if mode&011 != 0 {
			fmt.Println(fullPath)
			// if want to stop after finding the first occurence of desired exe
			return
		}
	}
}