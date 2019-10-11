/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	mathRand "math/rand"
	"time"
)


func generatePassword(length int) (string, error) {
	byteLength := length * 3 / 4
	if byteLength == 0 {
		byteLength = 1
	}
	bs := make([]byte, byteLength)
	fmt.Println(byteLength)
	_, err := io.ReadFull(rand.Reader, bs)
	if err != nil {
		return "", err
	}
	passw := base64.StdEncoding.EncodeToString(bs)
	fmt.Println(len(passw))
	return passw[:length], nil
}

func main() {
	mathRand.Seed(time.Now().Unix())
	passw, err := generatePassword(50)
	if err != nil {
		panic(err)
	}
	fmt.Println(passw)
}
