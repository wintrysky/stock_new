package dto

type LoginParam struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Status string `json:"status"` //ok error
	Type string `json:"type"` // account
	CurrentAuthority string `json:"currentAuthority"` // admin user guest
}
