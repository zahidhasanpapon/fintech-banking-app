package helpers

import "golang.org/x/crypto/bcrypt"

// HandleErr checks if have error, if yes send the panic
func HandleErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

// HashedAndSalt ed passwords imporves passwords security
func HashAndSalt(pass []byte) string {
	hashed, err := bcrypt.GenerateFromPassword(pass, bcrypt.MaxCost)
	HandleErr(err)

	return string(hashed)
}
