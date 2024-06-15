// package main

// import (
// 	"fmt"
// 	"time"

// 	"golang.org/x/crypto/bcrypt"
// )

// func generateHash(plainPassword string) {

// }

// func main() {
// 	// Prompt the user to enter a string
// 	fmt.Print("\nEnter a password :\n(This is for demonstration purposes)\n(please do not enter a password you actually use):")
// 	var ogPassword string
// 	fmt.Scanln(&ogPassword)

// 	// Generate a bcrypt hash for a password
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(ogPassword), bcrypt.DefaultCost)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	time.Sleep(time.Second)

// 	fmt.Println("\nHashed password:", string(hashedPassword))

// 	// Prompt the user to enter a string
// 	fmt.Print("\nRe-enter password, or enter incorrect password: ")
// 	var input string
// 	fmt.Scanln(&input)

// 	// Verify the password against the hashed value
// 	inputPassword := "mysecretpassword"
// 	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(inputPassword))
// 	if err != nil {
// 		fmt.Println("Password does not match")
// 		return
// 	}
// 	fmt.Println("Password matches")
// }
