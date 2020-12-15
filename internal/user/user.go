package user

import (
	"crypto/sha256"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	// "GOld/user"
)

const (
	layoutISO = "2006-01-02"      //1999-12-31 00:00:00 +0000 UTC func Parse(layout, value string) (Time, error)
	layoutUS  = "January 2, 2006" // December 31, 1999 func (t Time) Format(layout string) string
)

/*
date := "1999-12-31"
t, _ := time.Parse(layoutISO, date)
fmt.Println(t)                  // 1999-12-31 00:00:00 +0000 UTC
fmt.Println(t.Format(layoutUS)) // December 31, 1999
*/

type User struct {
	// Identity
	Internal_Username string `json:"internal_username"`
	Public_Username   string `json:"public_username"`
	// Username [32]byte `json:"username"`
	Name    string  `json:"name"`
	Surname string  `json:"surname"`
	ID      int64   `json:"id"`
	IP      string  `json:"IP"`
	Gas     float64 `json:"gas"`
}

func (u *User) serialize() []byte {
	marshalled, err := json.Marshal(u)
	if err != nil {
		fmt.Println("Marshalling Registered User failed.")
	}
	return marshalled
}

func (u *User) unique_username() string {
	to_sum := u.Name + u.Surname + strconv.Itoa(int(u.ID))
	// if u.public_key != nil {
	// 	to_sum += u.public_key
	// }
	sum := sha256.Sum256([]byte(to_sum))
	// return string(sum[:])
	b_64 := b64.StdEncoding.EncodeToString(sum[:])
	return b_64
}

func (u *User) inject() {
	// cur_time := time.Now() // 2009-11-10 23:00:00 +0000 UTC m=+0.000000000
	marshalled := u.serialize()
	fmt.Printf("[%T] %s", marshalled, marshalled)
	// log.Print(marshalled)

}

// func main() {
// 	u := User{
// 		Name:    "Cem",
// 		Surname: "Tunaboylu",
// 		ID:      1,
// 	}
// 	u.Internal_Username = u.unique_username()
// 	u.inject()
// }
