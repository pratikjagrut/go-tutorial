//The filepath package provides functions to parse and construct file paths
//in a way that is portable between operating systems
package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	//Join creates filepath,
	//takes any number of arg
	//and creates hierarchical path accordingly
	path := filepath.Join("/home", "userDir", "filename")
	fmt.Println("path:", path)

	//Join will also normalize paths by
	//removing superfluous separators and directory changes.
	fmt.Println(filepath.Join("/home//", "filename"))
	fmt.Println(filepath.Join("/home/../home", "filename"))

	//Dir() and Base() can split paths of Dir and file
	fmt.Println("Dir(path):", filepath.Dir(path))
	fmt.Println("Base(path):", filepath.Base(path))

	//Checks whether pach is absolute or not
	fmt.Println(filepath.IsAbs("Downloads/file"))  //false
	fmt.Println(filepath.IsAbs("/Downloads/file")) //true

	filename := "deployment.yaml"

	ext := filepath.Ext(filename) //returns extension of file
	fmt.Println(ext)

	onlyFilename := strings.TrimSuffix(filename, ext) //Trims extension
	fmt.Println(onlyFilename)

	//Finds relative path between basepath and targetpath
	rel, err := filepath.Rel("/home/temp", "/home/temp/work/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
	rel, err = filepath.Rel("/etc", "/home/userDir/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
