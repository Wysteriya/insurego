package models

type ReceiveRegisterIns struct {
	UserId string `json:"user_id"`
	PolicyRefId string `json:"policy_ref_id"`
	PolicyId string `json:"policy_id"`
	ClaimAmount string `json:"claim_amount"`
}

type SendRegisterIns struct {
	PolicyName string `json:"policy_name"`
	Policyid string `json:"policy_id"`
}
