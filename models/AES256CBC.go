package models

type AES256CBC struct {
	Key        string `json:"key"`
	IV         string `json:"iv"`
	PlainText  string `json:"plaintext"`
	CypherText string `json:"cyphertext"`
}
