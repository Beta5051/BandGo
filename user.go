package BandGo

type GetProfileResponse Profile

type GetPermissionsResponse struct {
	Permissions []string `json:"permissions"`
}

type Profile struct{
	UserKey string `json:"user_key"`
	ProfileImageUrl string `json:"profile_image_url"`
	Name string `json:"name"`
	IsAppMember bool `json:"is_app_member"`
	MessageAllowed bool `json:"message_allowed"`
}

type Author struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Role string `json:"role"`
	ProfileImageUrl string `json:"profile_image_url"`
	UserKey string `json:"user_key"`
}

func (c *Client) GetProfile(band_key string) (profile Profile, err error) {
	resp, err := c.request("GET", getProfileUrl, map[string]string{
		"band_key": band_key,
	})
	if err != nil {
		return
	}
	response := GetProfileResponse{}
	resp.dataConvert(&response)
	profile = Profile(response)
	return
}

func (c *Client) GetPermissions(band_key string, permissions []string) (return_permissions []string, err error){
	var permissions_list string
	for _, permission := range permissions {
		if permissions_list != "" {
			permissions_list += ","
		}
		permissions_list += permission
	}
	resp, err := c.request("GET", getPermissionsUrl, map[string]string{
		"band_key": band_key,
		"permissions": permissions_list,
	})
	response := GetPermissionsResponse{}
	resp.dataConvert(&response)
	return_permissions = response.Permissions
	return
}