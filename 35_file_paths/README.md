# Filepath

Package `filepath` implements utility routines for manipulating filename paths in a way compatible with the target operating system-defined file paths.

The filepath package uses either forward slashes or backslashes, depending on the operating system. To process paths such as URLs that always use forward slashes regardless of the operating system, see the path package.

## Import package

```go
import "filepath"
```

## Create file path which is portable over different OS

`Join` joins any number of path elements into a single path, separating them with an OS-specific Separator. Empty elements are ignored. The result is Cleaned. However, if the argument list is empty or all its elements are empty, Join returns an empty string. On Windows, the result will only be a UNC path if the first non-empty element is a UNC path.

Signature : `func Join(elem ...string) string`

```go
path := filepath.Join("/home", "userDir", "filename")
fmt.Println("path:", path)
```

Output: 

```
path: /home/userDir/filename
```

## Get directory and file or last element of the path

`Dir` returns all but the last element of the path, typically the path's directory.

signature: `func Dir(path string) string`

`Base` returns the last element of the path. Trailing path separators are removed before extracting the last element. If the path is empty, Base returns `.`.

signature: `func Base(path string) string`

```go
path := "/home/userDir/filename"
fmt.Println("Dir(path):", filepath.Dir(path))
fmt.Println("Base(path):", filepath.Base(path))
```

Output:

```
Dir(path): /home/userDir
Base(path): filename
```

## Splitting paths

`Split` splits path immediately following the final Separator, separating it into a directory and file name component. If there is no Separator in path, Split returns an empty dir and file set to path. The returned values have the property that `path = dir+file`.

Signature: `func Split(path string) (dir, file string)`

```go
path := "/home/userDir/filename"
dir, file := filepath.Split(path)
fmt.Println("Dir: ", dir)
fmt.Println("File: ", file)
```

Output: 

```
Dir:  /home/userDir/
File:  filename
```

***SplitList***

`SplitList` splits a list of paths joined by the OS-specific ListSeparator, usually found in PATH or GOPATH environment variables. Unlike strings.Split, SplitList returns an empty slice when passed an empty string.

Signature: `func SplitList(path string) []string`

```go
fmt.Println("SplitList: ", filepath.SplitList(os.Getenv("PATH")))
```

Output: 

```
SplitList: [/usr/lib/golang/bin /opt/pact/bin /usr/share/Modules/bin /usr/local/bin /usr/local/sbin /usr/bin /usr/sbin]
```

## Know whether the path is absolute or not

`IsAbs` reports whether the path is absolute.

Signature: `func IsAbs(path string) bool`

```go
fmt.Println("Is 'Downloads/file' absolute: ", filepath.IsAbs("Downloads/file"))
fmt.Println("Is '/Downloads/file' absolute: ", filepath.IsAbs("/Downloads/file"))
```

Output:

```
Is 'Downloads/file' absolute:  false
Is '/Downloads/file' absolute:  true
```

***You can refer main.go file for more examples***

To run:
```
go run main.go
```