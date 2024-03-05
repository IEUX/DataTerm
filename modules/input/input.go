package input

import (
	"fmt"
	"log"
	"syscall"

	"golang.org/x/term"
)

func GetPassword() string {
	fmt.Print("Enter host password: ")
	password, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatal("[INIT] Fail to read user input for password")
	}
	fmt.Print("\n")
	return string(password)
}
