package models


//response will be list of this struct
type SendTopIns struct {
	PolicyRefId string `json:"policy_ref_id"`
	InsuranceName string `json:"insurance_name"`
	CompanyName string `json:"company_name"`
	Type string `json:"type"`
	CoverageValue string `json:"coverage_value"`
	MonthlyCost string `json:"monthly_cost"`
	Description string `json:"description"`
	OtherDetails string `json:"other_details"`
	ClaimSuccessPercentage string `json:"claim_success_percentage"`
	LaunchDate string `json:"launch_date"`
	Claimed string `json:"claimed"`
	ClaimAmount string `json:"claim_amount"`
	ClaimDate string `json:"claim_date"`
	BuyDate string `json:"buy_date"`
	OtherDetails string `json:"other_details"`
}
