package models

type ReceiveClaimIns struct {
	PrivateKey  string `json:"private_key"`
	PublicKey   string `json:"public_key"`
	PolicyRefId string `json:"policy_ref_id"`
	ClaimAmount string `json:"claim_amount"`
	ClaimDate   string `json:"claim_date"`
}
