package models

type RSA struct {
	PlainText  int   `json:"plaintext"`
	CypherText int64 `json:"cyphertext"`
	DecryptedC int64 `json:"decryptedc"`
}
