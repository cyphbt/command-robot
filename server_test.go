package main

import (
	"log"
	"testing"
)

func TestVerifySignature(t *testing.T) {
	secretFail := "hello, world"
	secretOk := "It's a Secret to Everybody"
	payload := "Hello, World!"
	signature := "sha256=757107ea0eb2509fc211221cce984b8a37570b6d7586c22c46f4379c8b043e17"

	log.Println(verifySignature(secretFail, signature, []byte(payload)))
	log.Println(verifySignature(secretOk, signature, []byte(payload)))
}
