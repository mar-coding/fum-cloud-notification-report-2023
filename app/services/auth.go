package services

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/mar-coding/fum-cloud-notification-report-2023/app/models"
	"github.com/mar-coding/fum-cloud-notification-report-2023/app/utils"
)

func ValidateToken(signedToken string) (err error) {
	client := http.Client{}
	req, err := http.NewRequest("POST", "http://127.0.0.1:8082/user/validate", nil)
	if err != nil {
		return
	}

	req.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {signedToken},
	}

	res, err := client.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()
	var result models.Response

	// response body is []byte
	body, err := ioutil.ReadAll(res.Body)

	// Parse []byte to go struct pointer
	if err := json.Unmarshal(body, &result); err != nil {
		return err
	}
	if result.Id == 0 {
		return &utils.RequestError{
			StatusCode: 401,
			Err:        errors.New("Unauthorized"),
		}
	}
	return
}
