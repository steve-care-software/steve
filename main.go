package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"fmt"
)

func main() {
	// Generate an Ed25519 private and public key pair
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		fmt.Println("Error generating keys:", err)
		return
	}

	// The message we want to sign
	message := []byte("This is a secret message")

	// Sign the message with the private key
	signature := ed25519.Sign(privateKey, message)

	// Print the keys, signature, and the message
	fmt.Printf("Public Key: %x\n", publicKey)
	fmt.Printf("Private Key: %x\n", privateKey)
	fmt.Printf("Signature: %x\n", signature)
	fmt.Printf("Message: %s\n", message)

	// Verify the signature with the public key
	isValid := ed25519.Verify(publicKey, message, signature)
	if isValid {
		fmt.Println("Signature is valid!")
	} else {
		fmt.Println("Signature is not valid!")
	}
}
