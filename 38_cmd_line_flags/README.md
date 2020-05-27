# Command-line flags

Command-line flags are a common way to specify options for command-line programs. For example, in `ls -a` the `-a is a command-line flag`.

Go provides a `flag package` supporting basic command-line flag parsing. We’ll use this package to implement our example command-line program.

## Constructing flags

**First declare flags**

Basic flag declarations are available for `string, integer, and boolean options(flagName, defaultValue, usage)`

```go
stringPtr := flag.String("word", "foo", "a string")
intPtr := flag.Int("number", 42, "an int")
boolPtr := flag.Bool("ok", false, "a bool")
```
This flag function `retruns pointer to the value` not the value itself.

*We can also declare option that uses an existing variable*

```go
var extVar string
flag.StringVar(&extVar, "extvar", "bar", "a string var")
```

**Second parse the declared flags**

After all, flags are declared, `flag.Parse()` is called to execute the command-line parsing.

```go
flag.Parse()
```

**Use those flags as you want**

```go
fmt.Println("Word:", *stringPtr)
fmt.Println("Number:", *intPtr)
fmt.Println("Ok:", *boolPtr)
fmt.Println("ExtVar:", extVar)
```

Tailing var or positional arg

```go
fmt.Println("Tail:", flag.Args())
```

## Different running scenarios of this program, refer main.go

Running `without flags` will print the default value of the flags.

```
go run main.go

Word: foo
Number: 42
Ok: false
ExtVar: bar
Tail: []
```

Running with `--help` flag will print default value and usage of flags.

```
go run main.go --help

Usage of /tmp/go-build733756446/b001/exe/main:
  -extvar string
        a string var (default "bar")
  -number int
        an int (default 42)
  -ok
        a bool
  -word string
        a string (default "foo")
exit status 2
```

Running with `all the flags` and with tailing var or positional args.

```
go run main.go -word=life -number=8 -ok=true -extvar=hello a1 b2 c3 d4

Word: life
Number: 8
Ok: true
ExtVar: hello
Tail: [a1 b2 c3 d4]
```

Note that the flag package requires `all flags to appear before positional arguments` (otherwise, the flags will be interpreted as positional arguments).

```
go run main.go -word=life -number=8 a1 b2 c3 d4 -ok=true

Word: life
Number: 8
Ok: false
ExtVar: bar
Tail: [a1 b2 c3 d4 -ok=true]
```

```
go run main.go -word=life a1 b2 c3 d4

Word: life
Number: 42
Ok: false
ExtVar: bar
Tail: [a1 b2 c3 d4]
```
