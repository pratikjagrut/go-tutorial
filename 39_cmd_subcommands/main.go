//The flag package lets us easily define simple subcommands that have their own flags.
//Like we can organize flags under subcommands
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// We are creating new subcommand
	send := flag.NewFlagSet("send", flag.ExitOnError)
	// flags of send subcommand
	name := send.String("name", "", "Recipient name")
	message := send.String("msg", "", "Message")

	block := flag.NewFlagSet("block", flag.ExitOnError)
	username := block.String("username", "", "User name to be blocked")

	chpass := flag.NewFlagSet("chpass", flag.ExitOnError)
	newPass := chpass.String("newPass", "", "Change password")

	// Check length of cmd-line arg
	if len(os.Args) == 1 {
		fmt.Println("Use any command given below: ")
		fmt.Println(" - send  Send messages to your contacts")
		fmt.Println(" - block To block the user")
		fmt.Println(" - chpass To change the password")
		return
	}

	// switch case for parsing flags of subcommands
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

	// Check if flags are parsed and impliment logic accordingly
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

}
