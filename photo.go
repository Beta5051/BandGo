package BandGo

import "fmt"

type GetAlbumsResponse struct {
	Albums []Album `json:"items"`
	Paging struct {
		NextParams NextParams `json:"next_params"`
	} `json:"paging"`
}

type GetPhotosResponse struct {
	Photos []Photo `json:"items"`
	Paging struct {
		NextParams NextParams `json:"next_params"`
	} `json:"paging"`
}

type Album struct {
	PhotoAlbumKey string `json:"photo_album_key"`
	Name string `json:"name"`
	PhotoCount uint `json:"photo_count"`
	CreatedAt uint `json:"created_at"`
	Author Author `json:"author"`
}

type Photo struct {
	Height uint `json:"height"`
	Width uint `json:"width"`
	CreatedAt string `json:"created_at"`
	Url string `json:"url"`
	Author Author `json:"author"`
	PhotoAlbumKey string `json:"photo_album_key"`
	PhotoKey string `json:"photo_key"`
	CommentCount uint `json:"comment_count"`
	EmotionCount uint `json:"emotion_count"`
	IsVideoThumbnail bool `json:"is_video_thumbnail"`
}

func (c *Client) GetAlbums(band_key string, next_params interface{}) (albums []Album, return_next_params interface{}, err error){
	data := map[string]string{
		"band_key": band_key,
	}
	if v, ok := next_params.(NextParams); ok{
		data["after"] = v.After
		data["limit"] = fmt.Sprint(v.Limit)
	}
	resp, err := c.request("GET", getAlbumsUrl, data)
	if err != nil {
		return
	}
	response := GetAlbumsResponse{}
	resp.dataConvert(&response)
	albums = response.Albums
	if response.Paging.NextParams.After != "" {
		return_next_params = response.Paging.NextParams
	}
	return
}

func (c *Client) GetPhotos(band_key, photo_album_key string, next_params interface{}) (photos []Photo, return_next_params interface{}, err error){
	data := map[string]string{
		"band_key": band_key,
		"photo_album_key": photo_album_key,
	}
	if v, ok := next_params.(NextParams); ok{
		data["after"] = v.After
		data["limit"] = fmt.Sprint(v.Limit)
	}
	resp, err := c.request("GET", getPhotosUrl, data)
	if err != nil {
		return
	}
	response := GetPhotosResponse{}
	resp.dataConvert(&response)
	photos = response.Photos
	if response.Paging.NextParams.After != "" {
		return_next_params = response.Paging.NextParams
	}
	return
}