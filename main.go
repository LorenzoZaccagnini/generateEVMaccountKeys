package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/dustinxie/ecc"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	p256k1 := ecc.P256k1()

	// Generate a private key
	priv, err := ecdsa.GenerateKey(p256k1, rand.Reader)

	if err != nil {
		panic(err)
	}

	privateKeyBytes := priv.D.Bytes()
	fmt.Printf("Private key: %x", hex.EncodeToString(privateKeyBytes))

	// Generate a public key
	pub := priv.PublicKey

	// print as hexadecimal
	publicKeyBytes := elliptic.Marshal(p256k1, pub.X, pub.Y)
	fmt.Printf("Public key: %x", hex.EncodeToString(publicKeyBytes))

	// hash the public key
	hashedPublicKey := crypto.Keccak256(publicKeyBytes[1:])

	// cut first 12 bytes
	address := hashedPublicKey[12:]
	fmt.Printf("Address: %x", hex.EncodeToString(address))

}
