package encryption

import "golang.org/x/crypto/bcrypt"

func Encrypt(text string) (string, error) {
	hashedString, err := bcrypt.GenerateFromPassword([]byte(text), 10)
	if err != nil {
		return "", err
	}

	return string(hashedString), nil
}

func CompareHashedStrings(notHashed string, hashed string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(notHashed)); err != nil {
		return err
	}
	return nil
}
