package models

import "baby-chain/tools/data"

type ReceiveStatus struct {
	UserId string `json:"user_id"`
}

type SendStatus struct {
	data.Array
}
