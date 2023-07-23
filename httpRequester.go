package token

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Send(c *gin.Context, url, method string, payload map[string]interface{}) (map[string]interface{}, error) {
	client := &http.Client{}
	var req *http.Request
	var err error
	if payload == nil {
		req, err = http.NewRequest(method, url, nil)
	} else {
		pl, _ := json.Marshal(payload)

		req, err = http.NewRequest(method, url, bytes.NewBuffer(pl))
	}

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", c.GetHeader("Authorization"))

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 && resp.StatusCode > 200 {
		return nil, errors.New("Something Went wrong")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data map[string]interface{}

	err = json.Unmarshal(body, &data)

	return data, nil
}
