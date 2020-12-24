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
	byte_data_first, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}

	fmt.Printf("\n %s", message)
	byte_data_second, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}
	d1 := string(byte_data_first)
	d2 := string(byte_data_second)
	if d1 != d2 {
		fmt.Println("Information you provided do not match. ")
		return "", fmt.Errorf("DoNotMatch")
	}

	return strings.TrimSpace(d1), nil
}
