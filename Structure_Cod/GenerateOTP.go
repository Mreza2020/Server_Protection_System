package Server_Protection_System

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// # GenerateOTP
//
// This package is responsible for generating random numbers.
//
// # Parameters:
//
// length int = Number value For example : 10
//
// Number_length = What number should the numbers be from? For example, the numbers used should be from 1 to 5.
//
// # Returns:
//
// string = This is the output value that is returned, for example 123456.
func GenerateOTP(length int, Number_length int) string {
	otp := ""
	for i := 0; i < length; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(Number_length)))
		otp += fmt.Sprintf("%d", n.Int64())

	}
	return otp
}
