# EVER GREEN CLASSIC HELLO WORLD PROGRAM

*Here we will write hello world program in go.*

## First thing, package
In Go, source files are organized into system directories called packages, which enable code reusability across the Go applications.
When you build reusable pieces of code, you will develop a package as a shared library. But when you develop executable programs, you will use the package `main` for making the package as an executable program. The package `main` tells the Go compiler that the package should compile as an executable program instead of a shared library.

```go
package main
```

## Import
The keyword `import` is used for importing a package into other packages. When you import packages, the Go compiler will look on the locations specified by the environment variable `GOROOT` and `GOPATH`.

```go
// Single pkg/lib import
import "fmt"
// Multiple pkg/lib import
import(
    "fmt"
    "os"
)
```

## Main function
The `main` function in the package `main` will be the entry point of our executable program. When you build shared libraries, you will not have any main package and main function in the package.

```go
func main(){}
```

## Program 
You can refer `main.go` file.

```go
package main

import "fmt"

func main(){
    fmt.Println("Hello world")
}
```

Save it as `main.go` or `any-name-you-prefer.go`

To run:
```
go run main.go
```