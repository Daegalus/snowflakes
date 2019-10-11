package main

import (
	"crypto/rand"
	"encoding/binary"
	"flag"
	"fmt"
	"strings"
	"time"

	"snowflakes/base58"
)

var getTimeHash *string

func main() {
	cmdArgs()
	if *getTimeHash != "" {
		time, err := GetTimeFromHash(*getTimeHash)
		if err != nil {
			fmt.Printf("Error getting time from token. %s", err)
		}
		fmt.Printf("%d\n", time.Unix())
		return
	}

	snowflake, err := Snowflake()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(snowflake)
}

func cmdArgs() {
	getTimeHash = flag.String("time", "", "gets the time from the hash")
	flag.Parse()
}

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString Generates a number of bytes equal to the length provided, then Base58 encodes it with a custom alphabet.
func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base58.Encode(b, base58.SnowflakeAlphabet), err
}

// GetTimeFromHash Gets the time the hash was generated from the hash.
func GetTimeFromHash(hash string) (time.Time, error) {
	timeStr := strings.Split(hash, ".")[0]
	bytes, err := base58.Decode(timeStr, base58.SnowflakeAlphabet)
	if err != nil {
		return time.Unix(0, 0), fmt.Errorf("could not decode the time from the hash. %s", err)
	}

	timeSec := int64(binary.BigEndian.Uint32(bytes))
	return time.Unix(timeSec, 0), nil
}

// Snowflake generates a snowflake hash
func Snowflake() (string, error) {
	now := int32(time.Now().Unix())

	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(now))

	tokenTime := base58.Encode(b, base58.SnowflakeAlphabet)
	token, err := GenerateRandomString(6)
	if err != nil {
		return "", fmt.Errorf("could not generate token. %s", err)
	}

	return fmt.Sprintf("%s.%s", tokenTime, token), nil
}
