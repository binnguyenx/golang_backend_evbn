package service

import (
	"math/big"
	"crypto/rand"
)

const chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// service.GenPassService
type GenPass interface { //contract
	//remember when implement method then put it in the interface
	GeneratePassword(length int) (string, error)
}

type genPassService struct { //implementation
}

//hide implementation of the interface
// NewGenPassService returns a new GenPass service
func NewGenPassService() GenPass { //factory function
	return &genPassService{}
}

func (s *genPassService) GeneratePassword(length int) (string, error) { //method
	password := make([]byte, length)
	for i := 0; i < length; i++ {
		//we use rand.Int to generate a random number between 0 and the length of the chars string
		//because if hacker knows the length of the password, he can try to guess the password by trying all the possible combinations
		//we use big.NewInt to convert the length of the chars string to a big integer
		//we use rand.Int to generate a random number between 0 and the length of the chars string
		//we use rand.Int to generate a random number between 0 and the length of the chars string
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			return "", err
		}
		password[i] = chars[n.Int64()]
	}
	return string(password), nil
}
