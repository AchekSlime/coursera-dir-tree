package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	ans, err := mf(path, "", printFiles)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(out, ans)
	if err != nil {
		return err
	}
	return err
}

func mf(path string, prefix string, printFiles bool) (string, error) {
	ans := ""
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return ans, err
	}
	numberOfFiles := getNumberOfFiles(files, printFiles)
	printNum := 1
	for _, file := range files {
		if file.Name() == "static" {
			ans += ""
		}
		indent := "├───"
		newPrefix := prefix + "│\t"
		if printNum == numberOfFiles {
			indent = "└───"
			newPrefix = prefix + "\t"
		}
		if file.IsDir() {
			ans += prefix + indent + file.Name() + "\n"
			subTree, err := mf(path+"/"+file.Name(), newPrefix, printFiles)
			if err != nil {
				return ans, err
			}
			ans += subTree
			printNum++
		} else if printFiles && string(file.Name()[0]) != "." {
			size := strconv.FormatInt(file.Size(), 10) + "b"
			if file.Size() == 0 {
				size = "empty"
			}

			ans += prefix + indent + file.Name() + " (" + size + ")\n"
			printNum++
		}
	}
	return ans, nil
}

func getNumberOfFiles(files []os.FileInfo, printFiles bool) int {
	i := 0
	for _, file := range files {
		if file.IsDir() {
			i++
		} else if printFiles && string(file.Name()[0]) != "."{
			i++
		}
	}
	return i
}
