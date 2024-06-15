package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Printf(`
__________                                               .___   ___ ___               .__    .__                
\______   \_____    ______ ________  _  _____________  __| _/  /   |   \_____    _____|  |__ |__| ____    ____  
 |     ___/\__  \  /  ___//  ___/\ \/ \/ /  _ \_  __ \/ __ |  /    ~    \__  \  /  ___/  |  \|  |/    \  / ___\ 
 |    |     / __ \_\___ \ \___ \  \     (  <_> )  | \/ /_/ |  \    Y    // __ \_\___ \|   Y  \  |   |  \/ /_/  >
 |____|    (____  /____  >____  >  \/\_/ \____/|__|  \____ |   \___|_  /(____  /____  >___|  /__|___|  /\___  / 
                \/     \/     \/                          \/         \/      \/     \/     \/        \//_____/  `)
	fmt.Print("This is a demo, please do not use real passwords.")
	var input string
	var verify string
	for {
		fmt.Print("\nCreate new password or enter 'exit' to quit program:\n")
		fmt.Scanln(&input)

		switch strings.ToLower(input) {
		case "":
			fmt.Println("No password entered")
		case "exit":
			fmt.Println("Thank you for using this program!")
			os.Exit(3)
		default:
			fmt.Println("\nYou have entered: " + input)
			fmt.Print("Is this correct? (y/n): ")
			fmt.Scanln(&verify)

			if strings.ToLower(verify) == "y" {
				test(input)
			} else if strings.ToLower(verify) != "n" {
				fmt.Println("Input not registered")
				fmt.Println("You have entered: " + input)
				fmt.Println("Is this correct? (y/n)")
			}
		}

	}
}

func test(password string) {
	fmt.Println("\nYour password is: " + password)
	var input string
	for {
		input = ""
		fmt.Println("\nPlease enter your password.\nOr don't, I'm not a cop.\nEnter 'change' to change password\nEnter 'exit' to close program.")
		fmt.Scanln(&input)

		// time.Sleep(time.Second)

		switch strings.ToLower(input) {
		case password:
			fmt.Println("\nPassword entered correctly.")
			fmt.Println("Hashcode: TODO: generate hash and display")
			time.Sleep(time.Second)
		case "change":
			break
		case "exit":
			os.Exit(3)
		default:
			fmt.Println("\nYour password: " + password)
			fmt.Println("Hashcode: TODO generate hash and display")
			fmt.Println("\nInput: " + input)
			fmt.Println("Hashcode: TODO generate hash and display")
			time.Sleep(time.Second)
		}
	}
}
