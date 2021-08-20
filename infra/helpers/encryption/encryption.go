package encryption

import "golang.org/x/crypto/bcrypt"

func Encrypt(text string) (string, error) {
	hashedString, err := bcrypt.GenerateFromPassword([]byte(text), 25)
	if err != nil {
		return "", err
	}

	return string(hashedString), nil
}
