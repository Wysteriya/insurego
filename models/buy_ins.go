package models

type ReceiveRegisterIns struct {
	UserId string `json:"user_id"`
	PolicyRefId string `json:"policy_ref_id"`
	PolicyId string `json:"policy_id"`
}

type SendRegisterIns struct {
	PolicyRefId string `json:"policy_ref_id"`
}
