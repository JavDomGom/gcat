package models

type AES256CTR struct {
	Name string `json:"name"`
	Key  []byte `json:"key"`
	IV   []byte `json:"iv"`
	In   []byte `json:"in"`
	Out  []byte `json:"out"`
}
