package main

import (
	"fmt"

	"github.com/CTunaboylu/WhiteFox/pkg/ui"
)

// "../../pkg/ui"

func main() {

	Create_New_PGP()

	// var e *openpgp.Entity
	// e, err := openpgp.NewEntity("itis", "test", "itis@itis3.com", nil)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// for s, ident := range e.Identities {
	// 	fmt.Printf("%T : %s  \n", s, ident.UserId.Id) //
	// 	fmt.Printf("%T : %v \n", e.PrimaryKey.KeyId, e.PrimaryKey.KeyId)
	// }

	// Sign all the identities
	// for _, id := range e.Identities {
	// 	err := id.SelfSignature.SignUserId(id.UserId.Id, e.PrimaryKey, e.PrivateKey, nil)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// }

	// w, err := armor.Encode(os.Stdout, openpgp.PublicKeyType, nil)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer w.Close()

	// e.Serialize(w)
}

func Create_New_PGP() /* *openpgp.Entity */ {
	var pwd string
	fmt.Print("First, ")
	pwd, err := ui.Get_Sensitive_Information("Enter the password that you want to protect your keys with: ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\nPassword: %s\n ", pwd)

}
