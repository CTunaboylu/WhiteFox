package tokenize

import (
	"GOld/cryptography/DB"
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

const (
	user        string = "USR_"
	transaction string = "TRA_"
)

var empty = struct{}{}
var forms map[string]struct{} = map[string]struct{}{user: empty, transaction: empty}

func Generate_Base64_ID(size int, prefix string) (string, error) {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	encoded := base64.URLEncoding.EncodeToString(b)
	if _, ok := forms[prefix]; ok {
		id := prefix + encoded
		if DB.Check_ID_Unique(id) {
			return id, nil
		}
	}
	return "", fmt.Errorf("Prefix is undefined")
}
