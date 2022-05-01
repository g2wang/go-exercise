package orgaccnt

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/g2wang/go-exercise/orgaccnt/config"
	"github.com/g2wang/go-exercise/orgaccnt/models"
)

func Create(accountData models.AccountData) (bool, error) {
	postBody, _ := json.Marshal(
		map[string]*models.AccountData{
			"data": &accountData,
		})
	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post(config.Cfg.RestURL, "application/vnd.api+json", responseBody)
	if err != nil {
		log.Fatalf("error: %v", err)
		return false, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error: %v", err)
		return false, err
	}
	log.Println("status: " + http.StatusText(resp.StatusCode))
	if resp.StatusCode != 201 {
		err := errors.New("create account failure: http status: " + http.StatusText(resp.StatusCode))
		log.Fatalf("error: %v", err)
		return false, err
	}
	sb := string(body)
	log.Printf("resp len: %d", len(sb))
	return true, nil
}

func Fetch(UUID string) *models.AccountData {
	resp, err := http.Get(config.Cfg.RestURL + "/" + UUID)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer resp.Body.Close()
	var accountData map[string]*models.AccountData
	if err := json.NewDecoder(resp.Body).Decode(
		&accountData); err != nil {
		log.Fatalf("error: %v", err)
	}
	return accountData["data"]
}

func Delete(UUID string, version int64) (bool, error) {
	req, err := http.NewRequest("DELETE", config.Cfg.RestURL+"/"+UUID+"?version="+strconv.FormatInt(version, 10), nil)
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
		err := errors.New("delete account failure: http status: " + http.StatusText(resp.StatusCode))
		log.Fatalf("error: %v", err)
		return false, err
	}
	defer resp.Body.Close()
	return true, nil
}
