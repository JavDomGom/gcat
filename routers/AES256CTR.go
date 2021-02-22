package routers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"

	"github.com/JavDomGom/gcat/models"
)

func AES256CTR(w http.ResponseWriter, r *http.Request) {
	plaintext := []byte(r.URL.Query().Get("plaintext"))
	key, _ := hex.DecodeString(r.URL.Query().Get("key"))

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	var aes_ctr = models.AES256CBC{
		Key:        hex.EncodeToString(key),
		IV:         hex.EncodeToString(iv),
		PlainText:  string(plaintext),
		CypherText: hex.EncodeToString(ciphertext[aes.BlockSize:]),
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(aes_ctr)
}
