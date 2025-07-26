package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"kisaanSathi/pkg/config"

	"github.com/gin-gonic/gin"
)

type aesCipher struct {
	SecretKey   []byte
	BlockSize   int
	IsIvPresent bool
}

type AesCipherGroup interface {
	AuthTokenDecryption(encryptedData string) (string, error)
	Encryption(plainText string) (string, error)
	Decryption(encryptedData string, staticIV string) (string, error)
	//PrepareAuthorizationPayload(data interface{}) (string, error)
}

func NewAesCipherService(secretKey string, isIvPresent bool) AesCipherGroup {
	return &aesCipher{
		SecretKey:   []byte(secretKey),
		BlockSize:   aes.BlockSize,
		IsIvPresent: isIvPresent,
	}
}

func (a *aesCipher) pad(data []byte, blocklen int) []byte {
	padding := blocklen - len(data)%blocklen
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

func (a *aesCipher) unPad(data []byte) []byte {
	unpadding := int(data[len(data)-1])
	return data[:(len(data) - unpadding)]
}

func (a *aesCipher) AuthTokenDecryption(encryptedData string) (string, error) {
	iv := make([]byte, 16)
	encryptedByteData := []byte(encryptedData)
	//Take out IV from encrypted access toke
	if a.IsIvPresent {
		ivPosition := len(encryptedData) - 16
		iv = encryptedByteData[ivPosition:]
		encryptedByteData = encryptedByteData[:ivPosition]
	}

	data, err := base64.StdEncoding.DecodeString(string(encryptedByteData))
	if err != nil {
		return "", fmt.Errorf("error in decoding the string, error: %v", err)
	}

	if len(data) < a.BlockSize {
		return "", fmt.Errorf("error in decoding the string, encrypted too short")
	}

	// CBC mode always works in whole blocks.
	if len(data)%a.BlockSize != 0 {
		return "", fmt.Errorf("error in decoding the string, ciphertext is not a multiple of the block size")
	}

	cipherBlock, err := aes.NewCipher(a.SecretKey)
	if err != nil {
		return "", fmt.Errorf("error in creating new cipher, Error: %v", err)
	}

	decryptor := cipher.NewCBCDecrypter(cipherBlock, iv)

	decryptedBytes := make([]byte, len(data))
	decryptor.CryptBlocks(decryptedBytes, []byte(data))

	decryptedBytes = a.unPad(decryptedBytes)

	return string(decryptedBytes), nil
}

func (a *aesCipher) Encryption(plainText string) (string, error) {
	// Create a new AES cipher block using the secret key provided in the struct.
	cipherBlock, err := aes.NewCipher(a.SecretKey)
	if err != nil {
		return "", fmt.Errorf("error in creating new cipher, Error %v", err)
	}

	// Retrieve the block size of the AES cipher (usually 16 bytes for AES).
	blockSize := cipherBlock.BlockSize()

	staticIV := config.GetConfig().GetString("aes.iv")

	iv := []byte(staticIV)

	if len(iv) != 16 {
		return "", fmt.Errorf("static IV must be exactly 16 bytes")
	}
	// Create a new CBC (Cipher Block Chaining) encrypter using the cipher block and the IV.
	// CBC mode requires the IV to chain encryption between blocks.
	encrypter := cipher.NewCBCEncrypter(cipherBlock, iv)

	// Convert the plaintext string into a byte slice to prepare it for encryption.
	plainTextBytes := []byte(plainText)

	// Pad the plaintext bytes to ensure the length is a multiple of the block size.
	// AES encryption requires data in fixed-size blocks, so padding fills the remaining space.
	plainTextBytes = a.pad(plainTextBytes, blockSize)

	// Create a byte slice to store the resulting ciphertext.
	// The length of the ciphertext will be equal to the length of the padded plaintext.
	cipherText := make([]byte, len(plainTextBytes))

	// Perform the encryption operation. The encrypter processes the padded plaintext
	// and writes the encrypted result to the cipherText slice.
	encrypter.CryptBlocks(cipherText, plainTextBytes)

	// Encode the ciphertext to a base64-encoded string for easy storage or transmission.
	// Base64 ensures the encrypted data can be safely represented as text.
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func (a *aesCipher) Decryption(encryptedData string, staticIV string) (string, error) {

	// Convert the static IV to a byte slice
	iv := []byte(staticIV)

	if len(iv) != 16 {
		return "", fmt.Errorf("static IV must be exactly 16 bytes")
	}

	// Decode the base64-encoded string
	data, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return "", fmt.Errorf("error in decoding the string, error: %v", err)
	}

	// Ensure the data contains at least IV + 1 block
	if len(data) < aes.BlockSize {
		return "", fmt.Errorf("error in decoding the string, encrypted data too short")
	}

	cipherText := data

	// Create AES cipher block
	cipherBlock, err := aes.NewCipher(a.SecretKey)
	if err != nil {
		return "", fmt.Errorf("error in creating new cipher, Error: %v", err)
	}

	// Decrypt the ciphertext
	decryptor := cipher.NewCBCDecrypter(cipherBlock, iv)
	decryptedBytes := make([]byte, len(cipherText))
	decryptor.CryptBlocks(decryptedBytes, cipherText)

	// Unpad the decrypted plaintext
	decryptedBytes = a.unPad(decryptedBytes)

	// Return the plaintext as a string
	return string(decryptedBytes), nil
}

func EncryptData(data interface{}) (string, error) {

	encryptionKey := config.GetConfig().GetString("aes.secretKey")

	cipher := NewAesCipherService(encryptionKey, true)

	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("error converting piiData into json, error: %v", err)
	}

	encryptedData, err := cipher.Encryption(string(jsonData)) //Encrypted Data
	if err != nil {
		return "", fmt.Errorf("error encrypting piiData details, error: %v", err)
	}

	return encryptedData, nil

}

func DecryptData(data string, c *gin.Context) (string, error) {

	encryptionKey := config.GetConfig().GetString("aes.secretKey")

	IV := c.GetHeader(config.IV256)

	if len(IV) != 16 {
		return "", fmt.Errorf("static IV must be exactly 16 bytes")
	}
	cipher := NewAesCipherService(encryptionKey, true)

	encryptedData, err := cipher.Decryption(data, IV) //Encrypted Data
	if err != nil {
		return "", fmt.Errorf("error encrypting piiData details, error: %v", err)
	}

	return encryptedData, nil

}
