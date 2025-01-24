package domain

type Token struct {
	Value   string `json:"token"`
	IsValid bool   `json:"is_valid"`
}
