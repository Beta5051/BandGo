package BandGo

import "fmt"

type GetCommentsResponse struct {
	Comments []Comment `json:"items"`
	Paging struct {
		NextParams NextParams `json:"next_params"`
	} `json:"paging"`
}

type CreateCommentResposne struct {
	Message string `json:"message"`
}

type RemoveCommentResponse struct {
	Message string `json:"message"`
}

type Comment struct {
	BandKey string `json:"band_key"`
	Author Author `json:"author"`
	PostKey string `json:"post_key"`
	CommentKey string `json:"comment_key"`
	Content string `json:"content"`
	EmotionCount uint `json:"emotion_count"`
	IsAudioIncluded bool `json:"is_audio_included"`
	CreatedAt uint `json:"created_at"`
	Photo struct {
		Url string `json:"url"`
		Height uint `json:"height"`
		Width uint `json:"width"`
	} `json:"photo"`
}

func (c *Client) GetComments(band_key, post_key string, next_params interface{}, sort bool) (comments []Comment, return_next_params interface{}, err error){
	data := map[string]string{
		"band_key": band_key,
		"post_key": post_key,
		"sort": "+created_at",
	}
	if !sort {
		data["sort"] = "-created_at"
	}
	if v, ok := next_params.(NextParams); ok{
		data["after"] = v.After
		data["limit"] = fmt.Sprint(v.Limit)
	}
	resp, err := c.request("GET", getCommentUrl, data)
	if err != nil {
		return
	}
	response := GetCommentsResponse{}
	resp.dataConvert(&response)
	comments = response.Comments
	if response.Paging.NextParams.After != "" {
		return_next_params = response.Paging.NextParams
	}
	return
}

func (c *Client) CreateComment(band_key, post_key, body string) (message string, err error){
	resp, err := c.request("POST", createComment, map[string]string{
		"band_key": band_key,
		"post_key": post_key,
		"body": body,
	})
	if err != nil {
		return
	}
	response := CreateCommentResposne{}
	resp.dataConvert(&response)
	message = response.Message
	return
}

func (c *Client) RemoveComment(band_key, post_key, comment_key string) (message string, err error){
	resp, err := c.request("POST", removeComment, map[string]string{
		"band_key": band_key,
		"post_key": post_key,
		"comment_key": comment_key,
	})
	if err != nil {
		return
	}
	response := RemoveCommentResponse{}
	resp.dataConvert(&response)
	message = response.Message
	return
}