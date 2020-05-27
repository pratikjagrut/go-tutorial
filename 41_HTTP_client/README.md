# HTTP Request

Here we’re going to make a simple http GET requests using the `net/http` built-in package.

The below program is pretty straight forward.

```go
// Issue an HTTP GET request to a server
resp, err := http.Get("https://golang.org/doc/")
if err != nil {
    panic(err)
}
defer resp.Body.Close() // Close Body at the end

// Print the HTTP response status.
fmt.Println("Response status:", resp.Status)

fmt.Println()

scanner := bufio.NewScanner(resp.Body)
for i := 0; scanner.Scan() && i < 10; i++ { // Prints 10 lines from resp.Body
    fmt.Println(scanner.Text())
}

if err := scanner.Err(); err != nil { // checks for error
    panic(err)
}
```

Output:

```
Response status: 200 OK

<!DOCTYPE html>
<html lang="en">
<meta charset="utf-8">
<meta name="description" content="Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.">
<meta name="viewport" content="width=device-width, initial-scale=1">
<meta name="theme-color" content="#00ADD8">

  <title>Documentation - The Go Programming Language</title>

<link href="https://fonts.googleapis.com/css?family=Work+Sans:600|Roboto:400,700" rel="stylesheet">

```

***You can refer main.go file for examples***

To run:
```
go run main.go
```