package utils

import "golang.org/x/crypto/bcrypt"

func NormalizePassword(p string) []byte {
	return []byte(p)
}

func GeneratePassword(p string) string {
	//Normalize password from string to byte
	bytePassword := NormalizePassword(p)

	// MinCost is just an integer constant provided by the bcrypt package
	// along with DefaultCost & MaxCost. The cost can be any value
	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.MinCost) //the MinCost (4).
	if err != nil {
		return err.Error()
	}
	return string(hash)
}
