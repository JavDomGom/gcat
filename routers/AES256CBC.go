package routers

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"encoding/json"
	"net/http"

	"github.com/JavDomGom/gcat/models"
)

func PKCS5Padding(ciphertext []byte, blockSize, after int) []byte {
	padding := (blockSize - len(ciphertext)%blockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(ciphertext, padtext...)
}

func AES256Encrypt(plaintext, key, iv string, blockSize int) string {
	bKey := []byte(key)
	bIV := []byte(iv)
	bPlaintext := PKCS5Padding([]byte(plaintext), blockSize, len(plaintext))
	block, _ := aes.NewCipher(bKey)
	ciphertext := make([]byte, len(bPlaintext))
	mode := cipher.NewCBCEncrypter(block, bIV)
	mode.CryptBlocks(ciphertext, bPlaintext)

	return hex.EncodeToString(ciphertext)
}

func AES256CBC(w http.ResponseWriter, r *http.Request) {
	plaintext := r.URL.Query().Get("plaintext")
	key := r.URL.Query().Get("key")
	iv := r.URL.Query().Get("iv")

	var aes_cbc = models.AES256CBC{
		Key:        key,
		IV:         iv,
		PlainText:  plaintext,
		CypherText: AES256Encrypt(plaintext, key, iv, aes.BlockSize),
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(aes_cbc)
}
