package models

type ReceiveBuyIns struct {
	PrivateKey  string `json:"private_key"`
	PublicKey   string `json:"public_key"`
	UserId      string `json:"user_id"`
	PolicyRefId string `json:"policy_ref_id"`
}
