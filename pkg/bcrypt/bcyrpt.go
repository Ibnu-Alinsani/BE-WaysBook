package bcrypt

import "golang.org/x/crypto/bcrypt"

func HashingPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

func CheckPasswordHash(hashedPassword, password  string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil
}