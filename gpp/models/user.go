package models

type ReceiveUser struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Gender string `json:"gender"`
	State string `json:"state"`
	City string `json:"city"`
}

type SendUser struct {
	UserId string `json:"user_id"`
}
