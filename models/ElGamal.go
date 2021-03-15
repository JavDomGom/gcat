package models

type ElGamal struct {
	Bob_p        int64 `json:"bob_p"`
	Bob_a        int64 `json:"bob_a"`
	Bob_PriK     int64 `json:"bob_PriK"`
	Bob_PubK     int64 `json:"bob_PubK"`
	SecretNumber int64 `json:"secret_number"`
	Alice_v      int64 `json:"alice_v"`
	Alice_N1     int64 `json:"alice_N1"`
	Alice_N2     int64 `json:"alice_N2"`
	Bob_N3       int64 `json:"bob_N3"`
	N_decrypted  int64 `json:"N_decrypted"`
}
