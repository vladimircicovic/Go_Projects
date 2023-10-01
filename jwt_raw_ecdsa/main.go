package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

func main() {
	key, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)

	pvt, pub, err := pemKeyPair(key)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(pvt), "\n", string(pub))

	claims := &jwt.MapClaims{
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
		"name": "test",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES384, claims)

	tokenString, err := token.SignedString(key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tokenString)
	var ecdsaKey *ecdsa.PublicKey
	if ecdsaKey, err = jwt.ParseECPublicKeyFromPEM(pub); err != nil {
		fmt.Println("Unable to parse ECDSA public key: %v", err)
	}
	tkn, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return ecdsaKey, nil
	})
	fmt.Println(tkn.Valid)

}

func pemKeyPair(key *ecdsa.PrivateKey) (privKeyPEM []byte, pubKeyPEM []byte, err error) {
	der, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		return nil, nil, err
	}

	privKeyPEM = pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: der,
	})

	der, err = x509.MarshalPKIXPublicKey(key.Public())
	if err != nil {
		return nil, nil, err
	}

	pubKeyPEM = pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: der,
	})

	return
}
