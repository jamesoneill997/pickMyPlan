package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

//HashPassword uses bcrypt to securely encrypt the password
func HashPassword(password string) (string, error) {
	//encrypt password
	bytPwd := []byte(password)
	hashedPassword, passErr := bcrypt.GenerateFromPassword(bytPwd, bcrypt.DefaultCost)

	//handle err
	if passErr != nil {
		return "Error", errors.New("Internal server error")
	}

	//set user password to encrypted password before storage
	password = string(hashedPassword)

	return password, nil
}
