package models

type AuthRequest struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type AuthResponse struct {
	Token      string `json:"token"`
	ExpiringIn string `json:"expiring_in"`
}
