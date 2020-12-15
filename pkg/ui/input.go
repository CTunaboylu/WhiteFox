package ui

import (
	"fmt"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

func Get_Pwd() []byte {
	fmt.Println("Enter the password: ")
	var pwd string
	_, err := fmt.Scan(&pwd)
	if err != nil {
		fmt.Println(err)
	}
	return []byte(pwd)
	// return pwd
}
func Get_Pwd_String() string {
	fmt.Println("Enter the password: ")
	var pwd string
	_, err := fmt.Scan(&pwd)
	if err != nil {
		fmt.Println(err)
	}
	return pwd
}

func Get_Sensitive_Information(message string) (string, error) {

	fmt.Print(message)
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}

	password := string(bytePassword)
	return strings.TrimSpace(password), nil
}
