package BandGo

type NextParams struct {
	After string `json:"after"`
	Limit uint `json:"limit"`
	BandKey string `json:"band_key"`
	AccessToken string `json:"access_token"`
}