package security

import (
	"golang.org/x/crypto/bcrypt"
)

//Hash genera un hash de un password
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

//VerifyPassword compara el hash con la contraseña
func VerifyPassword(hashedPass, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(password))
}
