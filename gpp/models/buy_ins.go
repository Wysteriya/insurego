package models

type ReceiveBuyIns struct {
	PrivateKey  string `json:"private_key"`
	PublicKey   string `json:"public_key"`
	PolicyRefId string `json:"policy_ref_id"`
}
