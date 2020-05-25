package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	getError := func(err error) {
		if err != nil {
			panic(err)
		}
	}
	// Create temp
	err := os.Mkdir("temp", 0755) // 0755 are linux permissions
	getError(err)

	// defer os.RemoveAll("temp") // Delete directory after use

	// Create nested directories
	err = os.MkdirAll("temp/parent/child", 0755)
	getError(err)

	// Read temp
	c, err := ioutil.ReadDir("temp/parent")
	getError(err)

	fmt.Println("List temp/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir()) // IsDir returns bool
	}

	// Change dir
	err = os.Chdir("temp/parent/child")
	getError(err)

	pwd, err := os.Getwd()
	getError(err)

	fmt.Println("Current Dir: ", pwd)

	// Come back to temp directory
	err = os.Chdir("../../..")
	getError(err)

	fmt.Println()
	fmt.Println("Creating Temp dir and files")

	f, err := ioutil.TempFile("temp", "file")
	getError(err)
	fmt.Println("Temp file name:", f.Name())

	// os.Remove(f.Name()) to delete file as cleanup
	// as we are removing everything above in above defer os.RemoveAll("temp")

	f1, err := ioutil.TempFile("temp/parent", "file")
	getError(err)
	fmt.Println("Temp file name:", f1.Name())

	f2, err := ioutil.TempFile("temp/parent/child", "file")
	getError(err)
	fmt.Println("Temp file name:", f2.Name())

	dname, err := ioutil.TempDir("temp/parent/child", "grandChild")
	getError(err)
	fmt.Println("Temp dir name:", dname)

	fmt.Println()

	// look is called for every file or directory found recursively by filepath.Walk.
	look := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			fmt.Println(" ", path, "<- Dir")
		} else {
			fmt.Println(" ", path, "<- file")
		}
		return nil
	}

	fmt.Println("Tree of temp directory")
	// Walk walks the file tree rooted at root,
	// calling walkFn for each file or directory in the tree, including root
	err = filepath.Walk("temp", look)
	getError(err)

}
