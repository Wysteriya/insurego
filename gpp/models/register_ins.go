package models

type ReceiveRegisterIns struct {
	InsuranceName string `json:"insurance_name"`
	CompanyName string `json:"company_name"`
	Type string `json:"type"`
	CoverageValue string `json:"coverage_value"`
	MonthlyCost string `json:"monthly_cost"`
	Description string `json:"description"`
	OtherDetails string `json:"other_details"`
	ClaimSuccessPercentage string `json:"claim_success_percentage"`
	LaunchDate string `json:"launch_date"`
}

type SendRegisterIns struct {
	PolicyRefId string `json:"policy_ref_id"`
}
