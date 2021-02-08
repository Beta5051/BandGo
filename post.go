package BandGo

import "fmt"

type GetPostsResponse struct {
	Posts []Post `json:"items"`
	Paging struct {
		NextParams NextParams `json:"next_params"`
	} `json:"paging"`
}

type GetPostResponse struct {
	Post Post `json:"post"`
}

type CreatePostResponse struct {
	BandKey string `json:"band_key"`
	PostKey string `json:"post_key"`
}

type RemovePostResponse struct {
	Message string `json:"message"`
}

type Post struct {
	Content string `json:"content"`
	Author Author`json:"author"`
	PostKey string `json:"post_key"`
	CommentCount uint `json:"comment_count"`
	CreateAt uint `json:"created_at"`
	Photos []Photo `json:"photos"`
	EmotionCount uint `json:"emotion_count"`
	LatestComments []struct{
		Body string `json:"body"`
		Author Author `json:"author"`
		CreatedAt string `json:"created_at"`
	} `json:"latest_comments"`
	BandKey string `json:"band_key"`
}

func (c *Client) GetPosts(band_key, locale string, next_params interface{}) (posts []Post, return_next_params interface{}, err error){
	data := map[string]string{
		"band_key": band_key,
		"locale": locale,
	}
	if v, ok := next_params.(NextParams); ok{
		data["after"] = v.After
		data["limit"] = fmt.Sprint(v.Limit)
	}
	resp, err := c.request("GET", getPostsUrl, data)
	if err != nil {
		return
	}
	response := GetPostsResponse{}
	resp.dataConvert(&response)
	posts = response.Posts
	if response.Paging.NextParams.After != "" {
		return_next_params = response.Paging.NextParams
	}
	return
}

func (c *Client) GetPost(band_key, post_key string) (post Post, err error){
	resp, err := c.request("GET", getPostUrl, map[string]string{
		"band_key": band_key,
		"post_key": post_key,
	})
	if err != nil {
		return
	}
	response := GetPostResponse{}
	resp.dataConvert(&response)
	post = response.Post
	return
}

func (c *Client) CreatePost(band_key, content string, do_push bool) (return_band_key, post_key string, err error){
	resp, err := c.request("POST", createPost, map[string]string{
		"band_key": band_key,
		"content": content,
		"do_push": fmt.Sprint(do_push),
	})
	if err != nil {
		return
	}
	response := &CreatePostResponse{}
	resp.dataConvert(&response)
	return_band_key = response.BandKey
	post_key = response.PostKey
	return
}

func (c *Client) RemovePost(band_key, post_key string) (message string, err error) {
	resp, err := c.request("POST", removePost, map[string]string{
		"band_key": band_key,
		"post_key": post_key,
	})
	if err != nil {
		return
	}
	response := &RemovePostResponse{}
	resp.dataConvert(&response)
	message = response.Message
	return
}