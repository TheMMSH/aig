package pkg

import (
	"crypto/rand"
	"math/big"
)

func generateRandomSequence(length int) ([]byte, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	charsetLength := big.NewInt(int64(len(charset)))

	randomString := make([]byte, length)
	for i := 0; i < length; i++ {
		randomIndex, err := rand.Int(rand.Reader, charsetLength)
		if err != nil {
			return nil, err
		}
		randomString[i] = charset[randomIndex.Int64()]
	}

	return randomString, nil
}
