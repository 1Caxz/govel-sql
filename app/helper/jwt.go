package helper

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"govel/app/exception"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

func MakeECDSAToken(c jwt.Claims, method jwt.SigningMethod) string {
	token := jwt.NewWithClaims(method, c)
	key := loadECPrivateKeyFromDisk(os.Getenv("PRIVATE_KEY_FILE"))
	signed, err := token.SignedString(key)
	exception.PanicIfNeeded(err)
	return signed
}

func MakeRSAToken(c jwt.Claims, method jwt.SigningMethod) string {
	token := jwt.NewWithClaims(method, c)
	key := loadRSAPrivateKeyFromDisk(os.Getenv("PRIVATE_KEY_FILE"))
	signed, err := token.SignedString(key)
	exception.PanicIfNeeded(err)
	return signed
}

func ParseECDSAToken(token string, method jwt.SigningMethod) *jwt.Token {
	key := loadECPublicKeyFromDisk(os.Getenv("PUBLIC_KEY_FILE"))
	jwtToken, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return key, nil
	})
	exception.PanicIfNeeded(err)
	return jwtToken
}

func ParseRSAToken(token string, method jwt.SigningMethod) *jwt.Token {
	key := loadRSAPublicKeyFromDisk(os.Getenv("PUBLIC_KEY_FILE"))
	jwtToken, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return key, nil
	})
	exception.PanicIfNeeded(err)
	return jwtToken
}

func loadRSAPrivateKeyFromDisk(location string) *rsa.PrivateKey {
	keyData, err := os.ReadFile(location)
	exception.PanicIfNeeded(err)
	key, err := jwt.ParseRSAPrivateKeyFromPEM(keyData)
	exception.PanicIfNeeded(err)
	return key
}

func loadRSAPublicKeyFromDisk(location string) *rsa.PublicKey {
	keyData, err := os.ReadFile(location)
	exception.PanicIfNeeded(err)
	key, err := jwt.ParseRSAPublicKeyFromPEM(keyData)
	exception.PanicIfNeeded(err)
	return key
}

func loadECPrivateKeyFromDisk(location string) *ecdsa.PrivateKey {
	keyData, err := os.ReadFile(location)
	exception.PanicIfNeeded(err)
	key, err := jwt.ParseECPrivateKeyFromPEM(keyData)
	exception.PanicIfNeeded(err)
	return key
}

func loadECPublicKeyFromDisk(location string) *ecdsa.PublicKey {
	keyData, err := os.ReadFile(location)
	exception.PanicIfNeeded(err)
	key, err := jwt.ParseECPublicKeyFromPEM(keyData)
	exception.PanicIfNeeded(err)
	return key
}
