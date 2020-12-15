package ui

import "fmt"

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
