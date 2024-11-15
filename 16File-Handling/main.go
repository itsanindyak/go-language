package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
)

func showError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	content := "// This is file for database connection."
	filePath := "./go.txt"

	var file *os.File
	var err1,err error
	var a fs.FileInfo

	a, err = os.Stat(filePath)


	if err != nil {
		file, err1 = os.Create(filePath)
		showError(err1)
	} else {
		fmt.Println("Filename : ",a.Name())
		file, err1 = os.OpenFile(filePath, os.O_RDWR, 0644)
		showError(err1)
	}
	defer file.Close()

	_,err2 := io.WriteString(file,content)
	showError(err2)

	readfile(filePath)



}

func readfile(filename string) {
	content, err := os.ReadFile(filename)

	showError(err)
	fmt.Println(string(content))
}
