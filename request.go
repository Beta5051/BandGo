package BandGo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type DefaultResponse struct {
	ResultCode int `json:"result_code"`
	ResultData interface{} `json:"result_data"`
}

func (c *Client) request(method, url string, data map[string]string) (response DefaultResponse, err error){
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return
	}

	q := req.URL.Query()
	q.Add("access_token", c.Token)
	for k, v := range data {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := c.http.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return
	}

	if c.Debug {
		fmt.Printf("[%s] %s\n", method, req.URL.String())
		fmt.Printf("[Response] %s\n", string(bytes))
	}

	if response.ResultCode != 1 {
		err = errors.New(response.ResultData.(map[string]interface{})["message"].(string))
	}

	return
}

func (r DefaultResponse) dataConvert(v interface{}) {
	bytes, _ := json.Marshal(r.ResultData)
	json.Unmarshal(bytes, v)
}