# Directories and files

**Create Directory**

```go
err := os.Mkdir("temp", 0755) // 0755 are linux permissions
if err != nil{
    panic(err)
}
```

**Remove Dir after use**

```go
os.RemoveAll("temp")
```

**Create nested directories**

```go
err = os.MkdirAll("temp/parent/child", 0755)
if err != nil{
    panic(err)
}
```

**Read Directory**

```go
c, err := ioutil.ReadDir("temp/parent")
if err != nil{
    panic(err)
}

fmt.Println("List temp/parent")
for _, entry := range c {
    fmt.Println(" ", entry.Name(), entry.IsDir()) // IsDir returns bool
}
```

Output:

```
List temp/parent
  child true
```
**Change Dir**

```go
err = os.Chdir("temp/parent/child")
if err != nil{
    panic(err)
}

pwd, err := os.Getwd()
if err != nil{
    panic(err)
}

fmt.Println("Curretn Dir: ", pwd)
```

Output:

```
Current Dir:  /home/marserover/go/src/github.com/pratikjagrut/go-tutorial/36_directories_and_files/temp/parent/child
```

**Creating Temp dir and files**
`
```go
f, err := ioutil.TempFile("temp", "file")
if err != nil{
    panic(err)
}
fmt.Println("Temp file name:", f.Name())

// os.Remove(f.Name()) to delete file as cleanup
// as we are removing everything above in above defer os.RemoveAll("temp")

f1, err := ioutil.TempFile("temp/parent", "file")
if err != nil{
    panic(err)
}
fmt.Println("Temp file name:", f1.Name())

f2, err := ioutil.TempFile("temp/parent/child", "file")
if err != nil{
    panic(err)
}
fmt.Println("Temp file name:", f2.Name())

dname, err := ioutil.TempDir("temp/parent/child", "grandChild")
if err != nil{
    panic(err)
}
fmt.Println("Temp dir name:", dname)
```

Output:

```
Creating Temp dir and files
Temp file name: temp/file902878642
Temp file name: temp/parent/file402568553
Temp file name: temp/parent/child/file356467124
Temp dir name: temp/parent/child/grandChild016505219
```

**Tree of temp directory**

```go
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

// Walk walks the file tree rooted at root,
// calling walkFn for each file or directory in the tree, including root
err = filepath.Walk("temp", look)
if err != nil{
    panic(err)
}
```

Output:

```
Tree of temp directory
  temp <- Dir
  temp/file902878642 <- file
  temp/parent <- Dir
  temp/parent/child <- Dir
  temp/parent/child/file356467124 <- file
  temp/parent/child/grandChild016505219 <- Dir
  temp/parent/file402568553 <- file
```

***You can refer main.go file for examples***

To run:
```
go run main.go
```