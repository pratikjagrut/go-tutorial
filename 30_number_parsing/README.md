# Numbers Parsing

It is a very basic thing to parse a number from string in programming. GO has library `strconv` which provide rich functions to parse numbers or convert them from string to whichever number format you choose.

***There are many utilitarian functions available in `strconv` package below are some examples***

The most common numeric conversions are Atoi (string to int).

```go
i, _ := strconv.Atoi("786")
fmt.Println(i)
```

Output: 
```
786
```
Atoi is a convenience function for basic base-10 int parsing.

If you pass an arg other than int, Atoi will throw an error.

```go
_, e := strconv.Atoi("big")
fmt.Println(e)
```

Output: 

```
strconv.Atoi: parsing "big": invalid syntax
```

ParseInt, the 0 means infer the base from the string.
64 requires that the result fit in 64 bits.

```go
j, _ := strconv.ParseInt("738", 0, 64)
fmt.Println(j)
```

Output: 

```
738
```

ParseInt will recognize hex-formatted numbers.

```go
k, _ := strconv.ParseInt("0x3f5", 0, 64)
fmt.Println(k)
```

Output: 

```
1013
```

```go
u, _ := strconv.ParseUint("648", 0, 64)
fmt.Println(u)
```
Output: 

```
648
```

ParseFloat, this 5 tells how many bits of precision to parse.

```go
f, _ := strconv.ParseFloat("4.7494", 64)
fmt.Println(f)
```

Output: 

```
4.7494
```

***You can refer main.go file for examples***

To run:
```
go run main.go
```