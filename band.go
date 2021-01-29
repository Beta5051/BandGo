package BandGo

type GetBandsResponse struct{
	Bands []Band `json:"bands"`
}

type Band struct {
	Name string `json:"name"`
	BandKey string `json:"band_key"`
	Cover string `json:"cover"`
	MemberCount uint `json:"member_count"`
}

func (c *Client) GetBands() (bands []Band, err error) {
	resp, err := c.request("GET", getBandsUrl, nil)
	if err != nil {
		return
	}
	response := GetBandsResponse{}
	resp.dataConvert(&response)
	bands = response.Bands
	return
}