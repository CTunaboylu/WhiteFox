package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func hash_and_salt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	return string(hash)
}

func compare_passwords(hashed_pwd string, plain_pwd []byte) bool {
	byte_hash := []byte(hashed_pwd)
	err := bcrypt.CompareHashAndPassword(byte_hash, plain_pwd)
	if err != nil {
		fmt.Println(err)
	}
	return true
}

func create_hash(hash_target string) (string, error) {
	if len(hash_target) == 0 {
		return "", errors.New("Hash target is empty")
	}
	hasher := sha256.New()
	hasher.Write([]byte(hash_target))

	hash_sum := hasher.Sum(nil)
	str_2_hash := hex.EncodeToString(hash_sum)
	return str_2_hash, nil

}

func encrypt_AES(data []byte, pass_phrase []byte) ([]byte, error) {
	hash := hash_and_salt(pass_phrase)
	block, wtf := aes.NewCipher([]byte(hash))
	fmt.Printf("%v of type %T\n ", wtf, wtf)
	gcm, err := cipher.NewGCM(block) // Gallois counter mode
	if err != nil {
		fmt.Println(err.Error)
		return []byte(""), err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err.Error)
		return []byte(""), err
	}
	cipher := gcm.Seal(nonce, nonce, data, nil)
	return cipher, nil
}

func decrypt_AES(data []byte, pass_phrase []byte) ([]byte, error) {
	hash := hash_and_salt(pass_phrase)
	key := []byte(hash)
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err.Error)
		return []byte(""), err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println(err.Error)
		return []byte(""), err
	}
	nonce_Size := gcm.NonceSize()
	nonce, cipher := data[:nonce_Size], data[nonce_Size:]
	plain, err := gcm.Open(nil, nonce, cipher, nil)
	if err != nil {
		fmt.Println(err.Error)
		return []byte(""), err
	}
	return plain, nil
}

func encrypt_into_file_AES(filename string, data []byte, pass_phrase []byte) {
	f, _ := os.Create(filename)
	defer f.Close()
	cipher, _ := encrypt_AES(data, pass_phrase)
	f.Write(cipher)
}

func encrypt_file_AES(filename string, pass_phrase []byte, del_original bool) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	cipher, _ := encrypt_AES(data, pass_phrase)
	if del_original {
		f, err := os.Open(filename)
		defer f.Close()
		if err != nil {
			fmt.Println(err)
		}
		f.Write(cipher)
	} else {
		f, err := os.Create("encrypted.txt")
		defer f.Close()
		if err != nil {
			fmt.Println(err)
		}
		f.Write(cipher)
	}

}
