package goPandora

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

func encryptWithPublicKey(data, publicKeyPath string) (string, error) {
	publicKeyData, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		return "", fmt.Errorf("could not read public key file: %v", err)
	}
	block, _ := pem.Decode(publicKeyData)
	if block == nil {
		return "", fmt.Errorf("failed to decode PEM block containing public key")
	}
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("failed to parse public key: %v", err)
	}
	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("not an RSA public key")
	}

	hash := sha256.New()
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, rsaPublicKey, []byte(data), nil)
	if err != nil {
		return "", fmt.Errorf("encryption failed: %v", err)
	}

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func decryptWithPrivateKey(encryptedData, privateKeyPath string) (string, error) {
	privateKeyData, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		return "", fmt.Errorf("could not read private key file: %v", err)
	}
	block, _ := pem.Decode(privateKeyData)
	if block == nil {
		return "", fmt.Errorf("failed to decode PEM block containing private key")
	}
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("failed to parse private key: %v", err)
	}
	rsaPrivateKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		return "", fmt.Errorf("not an RSA private key")
	}
	ciphertext, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return "", fmt.Errorf("failed to decode base64 data: %v", err)
	}
	hash := sha256.New()
	plaintext, err := rsa.DecryptOAEP(hash, rand.Reader, rsaPrivateKey, ciphertext, nil)
	if err != nil {
		return "", fmt.Errorf("decryption failed: %v", err)
	}
	return string(plaintext), nil
}
