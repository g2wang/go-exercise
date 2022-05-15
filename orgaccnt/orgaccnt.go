package orgaccnt

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/g2wang/go-exercise/orgaccnt/models"
)

var URL string

func init() {
	URL = os.Getenv("ACCOUNT_SERVICE_URL")
	if len(URL) == 0 {
		URL = "http://localhost:8080/v1/organisation/accounts"
		log.Printf("ACCOUNT_SERVICE_URL environment var not set. Using default: %v", URL)
	}
}

func Create(accountData models.AccountData) (*models.AccountData, error) {
	postBody, _ := json.Marshal(
		map[string]*models.AccountData{
			"data": &accountData,
		})
	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post(URL, "application/vnd.api+json", responseBody)
	if err != nil {
		log.Fatalf("http post error: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	var accountDataResp map[string]*models.AccountData
	if err = json.NewDecoder(resp.Body).Decode(
		&accountDataResp); err != nil {
		log.Fatalf("error: %v", err)
		return nil, err
	}
	return accountDataResp["data"], nil
}

func Fetch(UUID string) (*models.AccountData, error) {
	resp, err := http.Get(URL + "/" + UUID)
	if err != nil {
		log.Fatalf("error: %v", err)
		return nil, err
	}
	defer resp.Body.Close()
	var accountData map[string]*models.AccountData
	if err = json.NewDecoder(resp.Body).Decode(
		&accountData); err != nil {
		log.Fatalf("error: %v", err)
		return nil, err
	}
	return accountData["data"], nil
}

func Delete(UUID string, version int64) (bool, error) {
	req, err := http.NewRequest("DELETE", URL+"/"+UUID+"?version="+strconv.FormatInt(version, 10), nil)
	if err != nil {
		log.Fatalf("error: %v", err)
		return false, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("error: %v", err)
		return false, err
	}
	log.Printf("delete status: %v", http.StatusText(resp.StatusCode))
	if resp.StatusCode != 204 {
		err = errors.New("delete account failure: http status: " + http.StatusText(resp.StatusCode))
		log.Fatalf("error: %v", err)
		return false, err
	}
	defer resp.Body.Close()
	return true, nil
}
