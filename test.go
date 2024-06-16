// package main

// import (
// 	"fmt"
// 	"time"

// 	"golang.org/x/crypto/bcrypt"
// )

// func generateHash(plainPassword string) {

// }

// func main() {
// 	fmt.Printf(`
//   ___ ___               .__    .__                 ________
//  /   |   \_____    _____|  |__ |__| ____    ____   \______ \   ____   _____   ____
// /    ~    \__  \  /  ___/  |  \|  |/    \  / ___\   |    |  \_/ __ \ /     \ /  _ \
// \    Y    // __ \_\___ \|   Y  \  |   |  \/ /_/  >  |    |   \  ___/|  Y Y  (  <_> )
//  \___|_  /(____  /____  >___|  /__|___|  /\___  /  /_______  /\___  >__|_|  /\____/
//        \/      \/     \/     \/        \//_____/           \/     \/      \/
// 					`)
// 	fmt.Print("(This is a demo, please do not use real passwords.)")
// 	var password string
// 	var verify string
// 	for {

// 		for {
// 			// Prompt the user to enter a string
// 			fmt.Print("\nEnter a password: ")
// 			fmt.Scanln(&password)

// 			verify = ""
// 			fmt.Print("\nVerify password: ")
// 			fmt.Scanln(&verify)

// 			if password == verify {
// 				break
// 			} else {
// 				fmt.Println("\nPasswords do not match.")
// 			}
// 		}

// 		// Generate a bcrypt hash for a password
// 		start := time.Now()
// 		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 		if err != nil {
// 			fmt.Println("Error:", err)
// 			return
// 		}
// 		t := time.Now()
// 		elapsed := t.Sub(start)

// 		fmt.Println("\nHashed password:", string(hashedPassword))
// 		fmt.Println("Time taken: " + elapsed.String())

// 		// Prompt the user to enter a string
// 		fmt.Print("\nRe-enter password, or enter incorrect password: ")
// 		var input string
// 		fmt.Scanln(&input)

// 		// Verify the password against the hashed value
// 		start = time.Now()
// 		err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(input))
// 		if err != nil {
// 			fmt.Println("Password does not match")
// 		} else {
// 			fmt.Println("Password matches")
// 		}
// 		t = time.Now()
// 		elapsed = t.Sub(start)

// 		fmt.Println("Time taken: " + elapsed.String())
// 	}
// }
