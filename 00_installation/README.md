# GOLANG

Go is an open-source, statically typed, compiled programming language built by Google.

It combines the best of both statically typed and dynamically typed languages and gives you the right mixture of efficiency and ease of programming. It is primarily suited for building fast, efficient, and reliable server-side applications.

Following are some of the most noted features of Go -

- Safety: Type safety and memory safety.
- Good support for Concurrency and communication.
- Efficient and latency-free Garbage Collection
- High-speed compilation
- Excellent Tooling support

## Installing GO

Go binary distributions are available for all major operating systems like Linux, Windows, and macOS. It’s super simple to install Go from the binary distributions.

If a binary distribution is not available for your operating system, [you can try installing Go from source](https://golang.org/doc/install/source).


### Linux

- Download the Linux distribution from Go’s official [download](https://golang.org/dl/) page.

    ```
    tar -xvf go1.14.3.linux-amd64.tar.gz
    ```

    It will create a directory named `go`.

- Move that dir to `/usr/local` where all other binaries reside.

    ```
    export PATH=$PATH:/usr/local/go/bin
    ```

- Custom Installation directory

    Instead of moving dir to `/usr/local` you can choose any other dir.

    ```
    mv go $HOME/
    ```

    Then set this custom location to `GOROOT` environment variable.

    ```
    export GOROOT=$HOME/go
    ```

    To make `GOROOT` permanent add it to `.bashrc` or `.zshrc` depending on which shell you are using or else your `GOROOT` will be unset once you end your terminal session or start a new terminal session.

### MacOS X

The simplest way to install GO in MacOS is Homebrew.

```
brew install go
```

### Windows

Download the Windows MSI installer file from Go’s official [download](https://golang.org/dl/) page. Open the installer and follow the on-screen instructions to install Go in your windows system. By default, the installer installs Go in `C:\Go`.

*Once installed try `go version` to check the installation.*

## Go Code organization

***Note: after the introduction of `Go modules in Go 1.11`, its no longer required to store Go code in the Go workspace. You can create your Go project in any directory outside of `GOPATH`. You can refer to [go docs on code organization](https://golang.org/doc/code.html). The following Go Code organization is still widely in use mostly because of its elegant organizing structure.***

### Workspace 

Go requires you to organize your code in a specific way -

By convention, all your Go code and the code must reside in a single workspace whose path is stored in the environment variable `GOPATH`.

The `workspace` dir is supposed to contain the following sub-dir:

- `src`: contains Go source files.

    The src directory typically contains many version control repositories containing one or more Go packages. Every Go source file belongs to a package. You generally create a new subdirectory inside your repository for every separate Go package. The tree for this looks like the following:

    ```
    go
    └── src
        └── github.com
            ├── pratikjagrut
            │   └── go-tutorial
            │       └── helloworld
            │           └── main.go
            └── user
                └── project
    ```

- `bin`: contains the executable binaries.

    The `Go tool` builds and installs executable binaries to this directory.

- `pkg`: contains Go package archives.

    All the non-executable packages (shared libraries) are stored in this directory. This is typically imported and used inside other executable packages.

### Setting GOPATH

#### Unix Systems (Linux and macOS)

```
mkdir $HOME/go_workspace
export GOPATH=$HOME/go_workspace
```

*To make `GOPATH` permanent add it to `.bashrc` or `.zshrc` depending on which shell you are using or else your `GOPATH` will be unset once you end your terminal session or start a new terminal session.*

#### Windows System

- Create the workspace folder at C:\go-workspace.

- Right-click on Start → click Control Panel → Select System and Security → click on System.

- From the menu on the left, select the Advanced system's settings.

- Click the Environment Variables button at the bottom.

- Click New from the User variables section.

- Type GOPATH into the Variable name field.

- Type C:\go-workspace into the Variable value field.

- Click OK.

***Note: GOPATH must be different than the path of your Go installation.***