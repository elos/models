package user

import "crypto/rand"

const potentialBytes = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyx_"

func RandomString(n int) string {

	if n <= 0 {
		return ""
	}

	var bytes = make([]byte, n)

	rand.Read(bytes)

	for i, b := range bytes {
		bytes[i] = potentialBytes[b%byte(len(potentialBytes))]
	}

	return string(bytes)
}

func NewKey() string {
	return RandomString(64)
}
