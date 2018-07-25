package encryption

import(
	"fmt"
	"golang.org/x/crypto/bcrpyt"
)

func BEncrypt(password string) string {
	//using a cost of 12 is sufficient for this project
	encrypted, err := bcrypt.GenerateFromPassword([]byte(password,12))
	return string(encrypted)
}

func BMatchingCheck(password, encrypted string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encrypted), []byte(password))
	return err == nil
}
