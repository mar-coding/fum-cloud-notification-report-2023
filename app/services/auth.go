package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/mar-coding/fum-cloud-notification-report-2023/app/models"
	"github.com/mar-coding/fum-cloud-notification-report-2023/app/utils"
)

func initAddress() string {
	err := godotenv.Load("./.env")

	if err != nil {
		log.Fatal("Failed to load the config")
	}
	var add string = os.Getenv("VALIDATE_USER_ADDRESS")
	conf := fmt.Sprintf("%s/user/validate", add)
	return conf
}

func ValidateToken(signedToken string) (int, error) {
	client := http.Client{}
	req, err := http.NewRequest("POST", initAddress(), nil)
	if err != nil {
		return 0, err
	}

	req.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {signedToken},
	}

	res, err := client.Do(req)
	if err != nil {
		return 0, err
	}

	defer res.Body.Close()
	var result models.Response

	// response body is []byte
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}

	// Parse []byte to go struct pointer
	if err := json.Unmarshal(body, &result); err != nil {
		return 0, err
	}
	id := result.Id
	if id == 0 {
		return 0, &utils.RequestError{
			StatusCode: 401,
			Err:        errors.New("Unauthorized"),
		}
	}
	return id, nil
}
