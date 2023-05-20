package models

import "baby-chain/tools/data"

type ReceiveINSHavs struct {
	PolicyRefId string `json:"policy_ref_id"`
}

type SendINSHavs struct {
	PubKeys data.Array `json:"pub_keys"`
}
