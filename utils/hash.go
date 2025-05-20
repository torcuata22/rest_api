package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hashedPassword), err
}

//bcrypt.GenerateFromPassword: This is the function that generates a hashed password from the original password.
//[]byte(password): We're converting the password string to a byte slice ([]byte) because bcrypt.GenerateFromPassword expects a byte slice as input.
//14: This is the cost parameter, which controls the complexity of the hash. Higher values make the hash more secure, but also slower to compute.

func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil //this will return false if password is valid, will NOT be nil if password invalid (bool)
}

//bcrypt has CompareHashAndPassword: This is the function that checks if the provided password matches the hashed password.
//[]byte(hashedPassword): We're converting the hashedPassword string to a byte slice ([]byte) because bcrypt.CompareHashAndPassword expects a byte slice as input.
//[]byte(password): We're converting the password string to a byte slice ([]byte) because bcrypt.CompareHashAndPassword expects a byte slice as input.
