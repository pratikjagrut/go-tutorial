# Testing

*What is a test?*

Putting it simply, `a test is a repeatable process that verifies whether or not something is working as intended`. 

NO programmers can write a bug-free program. This is the pivotal reason that testing is an essential part of the software development process.
Writing tests for our code is a good way to ensure quality and improve reliability.

GO provide very extensive support for testing unit components of our program.
A unit component can be functions, structs, methods and pretty much anything that can work standalone and contribute to a bigger goal.

## Testing in GO

Package `testing` provides support for automated testing of Go packages.
It is intended to be used in concert with the `go test` command,
which automates the execution of any function of the form `func TestXxx(*testing.T)`.

**Example**

Here is an example of a method we want to test in the `main` package.

```go
func sum(numbers ...int) int {
    sum := 0
    for _, n := range numbers {
        sum += n
    }
    return sum
}
```

We then write our test in a separate file. 
The test file can be in a different package or the same one.

*Note: If you are concerned about getting test function exported within the same package as of methods. Better to use different package for tests such as `package something_test`.*

Here we'll keep everything in `main` package.

**Characteristics of a Golang test function:**

- The first parameter of the testing function needs to be `t *testing.T`.

- Name of testing function begins with the word `Test` followed by a word or phrase `starting with a capital letter`
  `func TestSum(*testing.T)`.

- Calls `t.Error` or `t.Fail` to indicate a failure.

- `t.Log` can be used to provide non-failing debug information.

- Must be saved in a file named `something_test.go` such as: `sum_test.go`.

Below is simple test for sum function:

```go
func TestSum1(t *testing.T) {
    total := sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
    if total != 45 {
        t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 45)
    } else {
        t.Logf("(1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9), got: %d, want: %d.", total, 45)
    }
}
```

**Running tests**

There are two ways to run tests: 

The first way is `local directory mode` where we run the test using `go test` when our package is inside `$GOPATH`. 

```
go test

PASS
ok      std/github.com/pratikjagrut/go-tutorial/45_testing      0.002s
```

The second way is to run tests in the `package list mode`.
In this mode, we list what packages to test.

e.g.

- `go test .` to test package in the current directory.

- `go test package` when package belongs to `$GOPATH`. 
    
    We don't use this command inside a Go Module

- `go test ./test` to test package in `./test directory`.

- `go test ./...` to test all the package in the current directory.


**Running particular tests**

If we've lots of tests then we can choose particular tests to run using `-run` flag.

```
go test -v -run TestSum1

=== RUN   TestSum1
--- PASS: TestSum1 (0.00s)
    sum_test.go:12: (1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9), got: 45, want: 45.
PASS
ok      std/github.com/pratikjagrut/go-tutorial/45_testing      0.002s
```

**Running tests with verbosity**

We can print additional information about test function using verbose `-v` flag.

```
go test -v

=== RUN   TestSum1
--- PASS: TestSum1 (0.00s)
    sum_test.go:12: (1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9), got: 45, want: 45.
PASS
ok      std/github.com/pratikjagrut/go-tutorial/45_testing      0.002s
```

**Table-driven tests** 

In go, we often use table-driven testing to be able to test all function scenarios. Usually, the table is a slice of anonymous structs, however, you may define struct first or use an existing one. 

A series of related checks can be implemented by looping over a slice of test cases.

```go
func TestSum2(t *testing.T) {
    tables := []struct {
        x int
        y int
        n int
    }{
        {1, 1, 2},
        {1, 2, 3},
        {2, 2, 4},
        {5, 2, 7},
    }

    for _, table := range tables {
        total := sum(table.x, table.y)
        if total != table.n {
            t.Errorf("Sum of (%d+%d) was incorrect, got: %d, want: %d.",
                table.x, table.y, total, table.n)
        } else {
            t.Logf("(%d + %d) , got: %d, want: %d.", table.x, table.y, total, table.n)
        }
    }
}
```

This approach reduces the amount of repetitive code compared to repeating the same code for each test and makes it straightforward to add more test cases.

```
go test -v

=== RUN   TestSum2
--- PASS: TestSum2 (0.00s)
    sum_test.go:33: (1 + 1) , got: 2, want: 2.
    sum_test.go:33: (1 + 2) , got: 3, want: 3.
    sum_test.go:33: (2 + 2) , got: 4, want: 4.
    sum_test.go:33: (5 + 2) , got: 7, want: 7.
PASS
ok      std/github.com/pratikjagrut/go-tutorial/45_testing      0.002s
```

## Using Subtests

In `Go 1.7`, the testing package introduces a Run method on the `T` and `B` types that allows for the creation of `subtests` and `sub-benchmarks`. The introduction of subtests and sub-benchmarks enables better handling of failures, fine-grained control of which tests to run from the command line, control of parallelism, and often results in simpler and more maintainable code.

This test is a rewritten version of our earlier Table-driven tests using subtests:

```go
func TestSum3(t *testing.T) {
    tables := []struct {
        x int
        y int
        n int
    }{
        {1, 1, 2},
        {1, 2, 3},
        {2, 2, 4},
        {5, 2, 7},
    }

    for _, table := range tables {
        t.Run(fmt.Sprintf("Sum of %d, %d", table.x, table.y), func(t *testing.T) {
            total := sum(table.x, table.y)
            if total != table.n {
                t.Errorf("Sum of (%d+%d) was incorrect, got: %d, want: %d.",
                table.x, table.y, total, table.n)
            } else {
                t.Logf("(%d + %d) , got: %d, want: %d.", table.x, table.y, total, table.n)
            }
        })
    }
}
```

```
go test -v -run TestSum3

=== RUN   TestSum3
=== RUN   TestSum3/Sum_of_1,_1
=== RUN   TestSum3/Sum_of_1,_2
=== RUN   TestSum3/Sum_of_2,_2
=== RUN   TestSum3/Sum_of_5,_2
--- PASS: TestSum3 (0.00s)
    --- PASS: TestSum3/Sum_of_1,_1 (0.00s)
        sum_test.go:57: (1 + 1) , got: 2, want: 2.
    --- PASS: TestSum3/Sum_of_1,_2 (0.00s)
        sum_test.go:57: (1 + 2) , got: 3, want: 3.
    --- PASS: TestSum3/Sum_of_2,_2 (0.00s)
        sum_test.go:57: (2 + 2) , got: 4, want: 4.
    --- PASS: TestSum3/Sum_of_5,_2 (0.00s)
        sum_test.go:57: (5 + 2) , got: 7, want: 7.
PASS
ok      std/github.com/pratikjagrut/go-tutorial/45_testing      0.002s
```

An important benefit is if any tests case fails then it does not halt the execution it runs the next test case.

Alter the Table-driven test to fail.

```go
{
    {1, 1, 2},
    {1, 2, 6}, // Updated to fail
    {2, 2, 4},
    {5, 2, 7},
}
```

*Tests without subtests.*

This test is halted as soon as the second test case failed.

```
go test -v -run TestSum2

=== RUN   TestSum2
--- FAIL: TestSum2 (0.00s)
    sum_test.go:34: (1 + 1) , got: 2, want: 2.
    sum_test.go:31: Sum of (1+2) was incorrect, got: 3, want: 6.
    sum_test.go:32: Test case failed.
FAIL
exit status 1
FAIL    std/github.com/pratikjagrut/go-tutorial/45_testing      0.002s
```

*Tests with subtests*

Even though the second test case failed it continued to the next tests.

```
go test -v -run TestSum3

=== RUN   TestSum3
=== RUN   TestSum3/Sum_of_1,_1
=== RUN   TestSum3/Sum_of_1,_2
=== RUN   TestSum3/Sum_of_2,_2
=== RUN   TestSum3/Sum_of_5,_2
--- FAIL: TestSum3 (0.00s)
    --- PASS: TestSum3/Sum_of_1,_1 (0.00s)
        sum_test.go:58: (1 + 1) , got: 2, want: 2.
    --- FAIL: TestSum3/Sum_of_1,_2 (0.00s)
        sum_test.go:55: Sum of (1+2) was incorrect, got: 3, want: 6.
        sum_test.go:56: test case failed.
    --- PASS: TestSum3/Sum_of_2,_2 (0.00s)
        sum_test.go:58: (2 + 2) , got: 4, want: 4.
    --- PASS: TestSum3/Sum_of_5,_2 (0.00s)
        sum_test.go:58: (5 + 2) , got: 7, want: 7.
FAIL
exit status 1
FAIL    std/github.com/pratikjagrut/go-tutorial/45_testing      0.002s
```

**Setup and Tear-down**

Subtests and sub-benchmarks can be used to manage common setup and tear-down code:

```go
func TestFoo(t *testing.T) {
    // <setup code>
    t.Run("A=1", func(t *testing.T) { ... })
    t.Run("A=2", func(t *testing.T) { ... })
    t.Run("B=1", func(t *testing.T) {
        if !test(foo{B:1}) {
            t.Fail()
        }
    })
    // <tear-down code>
}
```

The setup and tear-down code will run if any of the enclosed subtests are run and will run at most once. This applies even if any of the subtests call Skip, Fail, or Fatal.


**Control of Parallelism**

Subtests allow fine-grained control over parallelism. To understand how to use subtests in the way it is important to understand the semantics of parallel tests.

Each test is associated with a test function. A test is called a parallel test if its test function calls the `Parallel method on its instance of testing.T`. A parallel test `never runs concurrently with a sequential test` and its execution is suspended until its calling test function, that of the parent test has returned. The `-parallel` flag defines the maximum number of parallel tests that can run in parallel.

`A test blocks until its test function returns and all of its subtests have completed`. This means that the parallel tests that are run by a sequential test will complete before any other consecutive sequential test is run.

This behaviour is identical for tests created by Run and top-level tests. In fact, under the hood top-level tests are implemented as subtests of a hidden master test.

```go
func TestGroupedParallel(t *testing.T) {
    for _, tc := range testCases {
        tc := tc // capture range variable
        t.Run(tc.Name, func(t *testing.T) {
            t.Parallel()
            if got := foo(tc.in); got != tc.out {
                t.Errorf("got %v; want %v", got, tc.out)
            }
            ...
        })
    }
}
```

The outer test will not complete until all parallel tests started by Run have completed. As a result, no other parallel tests can run in parallel to these parallel tests.

*Note that we need to `capture the range variable` to ensure that `tc` gets bound to the correct instance.*

The `TestSum3` is modified for parallelism.

```go
func TestSum4(t *testing.T) {
    tables := []struct {
        x int
        y int
        n int
    }{
        {1, 1, 2},
        {1, 2, 3},
        {2, 2, 4},
        {5, 2, 7},
    }

    for _, table := range tables {
        table := table // capture range variable
        t.Run(fmt.Sprintf("Sum of %d, %d", table.x, table.y), func(t *testing.T) {
            total := sum(table.x, table.y)
            t.Parallel()
            if total != table.n {
                t.Errorf("Sum of (%d+%d) was incorrect, got: %d, want: %d.", table.x, table.y, total, table.n)
                t.Fatal("test case failed.")
            } else {
                t.Logf("(%d + %d) , got: %d, want: %d.", table.x, table.y, total, table.n)
            }
        })
    }
}
```

```
go test -v -run TestSum3
=== RUN   TestSum3
=== RUN   TestSum3/Sum_of_1,_1
=== PAUSE TestSum3/Sum_of_1,_1
=== RUN   TestSum3/Sum_of_1,_2
=== PAUSE TestSum3/Sum_of_1,_2
=== RUN   TestSum3/Sum_of_2,_2
=== PAUSE TestSum3/Sum_of_2,_2
=== RUN   TestSum3/Sum_of_5,_2
=== PAUSE TestSum3/Sum_of_5,_2
=== CONT  TestSum3/Sum_of_1,_1
=== CONT  TestSum3/Sum_of_5,_2
=== CONT  TestSum3/Sum_of_2,_2
=== CONT  TestSum3/Sum_of_1,_2
--- PASS: TestSum3 (0.00s)
    --- PASS: TestSum3/Sum_of_1,_1 (0.00s)
        sum_test.go:60: (1 + 1) , got: 2, want: 2.
    --- PASS: TestSum3/Sum_of_5,_2 (0.00s)
        sum_test.go:60: (5 + 2) , got: 7, want: 7.
    --- PASS: TestSum3/Sum_of_2,_2 (0.00s)
        sum_test.go:60: (2 + 2) , got: 4, want: 4.
    --- PASS: TestSum3/Sum_of_1,_2 (0.00s)
        sum_test.go:60: (1 + 2) , got: 3, want: 3.
PASS
ok      std/github.com/pratikjagrut/go-tutorial/45_testing      0.002s
```

**Cleaning up after a group of parallel tests**

In the previous example we used the semantics to wait on a group of parallel tests to complete before commencing other tests. The same technique can be used to clean up after a group of parallel tests that share common resources:

```go
func TestTeardownParallel(t *testing.T) {
    // <setup code>
    // This Run will not return until its parallel subtests complete.
    t.Run("group", func(t *testing.T) {
        t.Run("Test1", parallelTest1)
        t.Run("Test2", parallelTest2)
        t.Run("Test3", parallelTest3)
    })
    // <tear-down code>
}
```

## Test coverage

Test Coverage is the percentage of your code covered by test suit. In layman’s language, it is the measurement of how many lines of code in your package were executed when you ran your test suit. Go provide built-in functionality to check your code coverage.

```
go test -cover

PASS
coverage: 70.0% of statements
ok      std/github.com/pratikjagrut/go-tutorial/45_testing      0.002s
```

***You can refer sum_test.go, greet_test.go file for examples***

To run:
```
go test -v
```