package models

import "baby-chain/tools/data"

type ReceiveRegisterIns struct {
	PrivateKey string    `json:"private_key"`
	PublicKey  string    `json:"public_key"`
	Data       data.Data `json:"data"`
}

type SendRegisterIns struct {
	PolicyRefId string `json:"policy_ref_id"`
}
