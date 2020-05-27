# Subcommands

Some command-line tools, like the go tool or git, have many subcommands, each with its own set of flags. For example, go build and go get are two different subcommands of the go tool. The `flag package` lets us easily define simple subcommands that have their flags.

## Constructing subcommands

**Declare subcommands**

We declare a subcommand using the `NewFlagSet` function and proceed to define new flags specific for this subcommand.

Declaring `send` subcommand which will send a message.

```go
send := flag.NewFlagSet("send", flag.ExitOnError)
```

Declaring flags for `send` subcommand.

```go
name := send.String("name", "", "Recipient name")
message := send.String("msg", "", "Message")
```

Similarly declare `block` and `chpass` subcommands and their flags.

```go
block := flag.NewFlagSet("block", flag.ExitOnError)
username := block.String("username", "", "User name to be blocked")

chpass := flag.NewFlagSet("chpass", flag.ExitOnError)
newPass := chpass.String("newPass", "", "Change password")
```

**Check length of cmd-line arg**

```go
if len(os.Args) == 1 {
    fmt.Println("Use any command given below: ")
    fmt.Println(" - send  Send messages to your contacts")
    fmt.Println(" - block To block ther user")
    fmt.Println(" - chpass To change the password")
    return
}
```

**switch case for parsing flags of subcommands**

```go
switch os.Args[1] {
case "send":
    send.Parse(os.Args[2:])
case "block":
    block.Parse(os.Args[2:])
case "chpass":
    chpass.Parse(os.Args[2:])
default:
    fmt.Println("Command not found: ", os.Args[1])
    os.Exit(2)
}
```

**Check if flags are parsed and implement logic accordingly**

```go
if send.Parsed() {
    if *name == "" {
        fmt.Println("Enter recipient name using recipient using -name option.")
        return
    }

    if *message == "" {
        fmt.Println("Enter message using -msg option.")
        return
    }

    fmt.Println("Sent to: \t", *name)
    fmt.Println("Message: \t", *message)
}

if block.Parsed() {
    if *username == "" {
        fmt.Println("Enter user name to block using -username option.")
        return
    }

    fmt.Printf("User %v blocked!\n", *username)
}

if chpass.Parsed() {
    if *newPass == "" {
        fmt.Println("Provide new password using -pass option.")
        return
    }

    fmt.Println("Password changed!")
}
```

## Different scenarios for running this program, refer main.go

```
go run main.go

Use any command given below: 
 - send  Send messages to your contacts
 - block To block the user
 - chpass To change the password
```

```
go run main.go send -name="James Bond" -msg="New Mission"

Sent to:     James Bond
Message:     New Mission
```

```
go run main.go block -username=Thanos

User Thanos blocked!
```

```
go run main.go chpass -newPass=1234567890

Password changed!
```