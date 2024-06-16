package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	fmt.Printf(`
  ___ ___               .__    .__                 ________
 /   |   \_____    _____|  |__ |__| ____    ____   \______ \   ____   _____   ____
/    ~    \__  \  /  ___/  |  \|  |/    \  / ___\   |    |  \_/ __ \ /     \ /  _ \
\    Y    // __ \_\___ \|   Y  \  |   |  \/ /_/  >  |    |   \  ___/|  Y Y  (  <_> )
 \___|_  /(____  /____  >___|  /__|___|  /\___  /  /_______  /\___  >__|_|  /\____/
       \/      \/     \/     \/        \//_____/           \/     \/      \/
					`)
	fmt.Println("(This is a demo, inputs are not private, enter 'exit' at any time to quit.)")

	for {
		password := newPassword()
		testPassword(password)
	}
}

func newPassword() string {
	for {
		password := userInput("Enter a password")
		if confirmPW := verify("Verify password"); confirmPW {
			return password
		}
	}
}

func generate(text string) []byte {
	start := time.Now()

	hashedText, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}

	end := time.Now()
	fmt.Printf("\nText: %s\nHash: %s\nTime: %s\n\n", text, string(hashedText), end.Sub(start).String())

	return hashedText
}

func testPassword(password string) {
	hashedPW := generate(password)

	for {
		input := userInput("Test password, or 'change' to change password")

		if strings.ToLower(input) == "change" {
			return
		} else {
			compare(input, password, hashedPW)
		}
	}
}

func compare(input, password string, hashedPW []byte) {
	start := time.Now()
	err := bcrypt.CompareHashAndPassword(hashedPW, []byte(input))
	if err == nil {
		end := time.Now()
		fmt.Println("Password matches")
		fmt.Printf("Time taken: %s\n", end.Sub(start).String())
	} else {
		fmt.Println("Password does not match")
		end := time.Now()
		fmt.Printf("Time taken: %s\n", end.Sub(start).String())
		time.Sleep(time.Second / 2)
		if showHash := verify("Compare saved hash with input hash?"); showHash {
			fmt.Printf("\nText: %s\nHash: %s", password, hashedPW)
			generate(input)
		} else {
			fmt.Println()
		}
	}
}

func verify(prompt string) bool {
	for {
		switch verify := userInput(prompt + " (y/n)"); verify {
		case "y":
			return true
		case "n":
			return false
		case "":
			fmt.Println("No input detected.")
			continue
		default:
			fmt.Println("Input not accepted")
			continue
		}
	}
}

func userInput(prompt string) string {
	var input string

	fmt.Printf("%s: ", prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	input = scanner.Text()

	if strings.ToLower(input) == "exit" {
		os.Exit(3)
	}

	return input
}
