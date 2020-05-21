# Defer, Panic and Recover

## Defer

A `defer statement` pushes a function call onto a list. 
The list of saved calls is executed after the surrounding function returns. 
`Defer is commonly used to simplify functions that perform various clean-up or closing actions.`

We can defer any function using `defer` keyword.

`defer cleanup()`

**The behaviour of defer statements is straightforward and predictable. There are three simple rules:**

- *A deferred function's arguments are evaluated when the defer statement is evaluated.*

    ```go
    func a() {
        i := 0
        defer fmt.Println(i)
        i++
        return
    }
    ```

    The above example will print `0`. Because `defer statement` will be evaluated after function returns.

- *Deferred function calls are executed in Last In First Out order after the surrounding function returns.*

    ```go
    func b() {
        for i := 0; i < 4; i++ {
            defer fmt.Printf(" %d", i)
        }
    }
    ```

    The above example will print 3 2 1 0.

- *Deferred functions may read and assign to the returning function's named return values.*

    ```go
    func c() (i int) {
        defer func() { i++ }()
        return 1
    }
    ```
    
    The above example has named return value `i`, even though the return statement `return 1`,
    `deferred function` will increment `i by 1` and function will return `2`.

## Panic

`Panic` is a built-in function that stops or aborts normal flow of execution if function or program returns something unexpected. Panic is a sort of exception.

When a function calls panic execution of that function is stoped any `deferred function` will work normally and function will return to its caller and that function will act as panic to the caller and this stacked up the process and program crashes.

Panics can be invoked by calling `panic(arg)`.

Panic is necessary if something unexpected happens.

```go
_, err := os.Open("file.txt")
if err != nil {
    panic(err)
}
```

The above example, if `os.Open(file.txt)` won't find any file of name `file.txt`, it will give some error value to `err` variable. If err is not equal to nil and will show value stored in `err`.

Below is the output of the above panic example.

```
panic: open file: no such file or directory

goroutine 1 [running]:
main.main()
    ./main.go:13 +0xd0
exit status 2
Process exiting with code: 1
```

We can print custom error message in panic.

`panic("Something went wrong!")`

## Recover

`Recover` is a built-in function that `regains control of a panicking goroutine`. During normal execution, a call to recover will return nil and have no other effect. If the current goroutine is panicking, a call to recover will capture the value given to panic and resume normal execution. `Recover is only useful inside deferred functions.`

The signature of recover is: `func recover() interface{} `

In below example, we are calling `firstFunc()` from `main`, `secondFunc(arg)` from `firstFunc()` and `secondFunc(arg)` will call itself `recursively` till we access the index of array `0 <= index < 5`, once the we try to access index 5 it will panic with message: `runtime error: index out of range [5] with length 5` and `seconfFunc will stop execution` and `will return to its caller firstFunc` where we have `deferred function` which collects value of `recover()` and check if its `nil` and normal execution will be regained hereafter. Function `firstFunc` will return to its caller which is `main` and main will execute normally.

```go
func main() {
    fmt.Println("Calling firstFunc from main")
    firstFunc()
    fmt.Println("Returned normally from firstFunc.")
}

func firstFunc() {
    defer func() {
        if err := recover(); err != nil {
            fmt.Println("Recovered in firstFunc: ", err)
        }
    }()

    fmt.Println("Calling secondFunc from firstFunc")
    secondFunc(0)
    fmt.Println("Returned normally from secondFunc.")
}

func secondFunc(i int) {
    var arr [5]int
    arr[i] = i + 1
    defer fmt.Printf("Defer -> Index: %d, Value: %d\n", i, arr[i])
    fmt.Printf("Index: %d, Value: %d\n", i, arr[i])
    secondFunc(i + 1)
}
```

Output: 

```
Calling firstFunc from main
Calling secondFunc from firstFunc
Index: 0, Value: 1
Index: 1, Value: 2
Defer -> Index: 1, Value: 2
Defer -> Index: 0, Value: 1
Recovered in firstFunc:  runtime error: index out of range [2] with length 2
Returned normally from firstFunc.
Cleaning function ...
Process exiting with code: 0
```

***Recover works only when it is called from the same goroutine. It's not possible to recover from a panic that has happened in a different goroutine.***

If we just run `secundFunc` as a separate goroutine then panic won't be handled by recover.

Just make the following changes to the `firstFunc`.

```go
go secondFunc(0)
time.Sleep(1 * time.Second
```

 `Sleep` in the next line is just to ensure that the program does not terminate before `secondFunc(arg)` has finished running.

 Output:

 ```
 Calling firstFunc from main
Calling secondFunc from firstFunc
Index: 0, Value: 1
Index: 1, Value: 2
Defer -> Index: 1, Value: 2
Defer -> Index: 0, Value: 1
panic: runtime error: index out of range [2] with length 2

goroutine 6 [running]:
main.secondFunc(0x2)
	.../go-tutorial/28_defer_panic_recover/main.go:32 +0x253
main.secondFunc(0x1)
	.../go-tutorial/28_defer_panic_recover/main.go:35 +0x21d
main.secondFunc(0x0)
	.../go-tutorial/28_defer_panic_recover/main.go:35 +0x21d
created by main.firstFunc
	.../go-tutorial/28_defer_panic_recover/main.go:25 +0xde
exit status 2
Process exiting with code: 1
 ```

 ***You can refer main.go file for examples***

To run:
```
go run main.go
```