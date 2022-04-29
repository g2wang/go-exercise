package orgaccnt

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/g2wang/go-exercise/orgaccnt/config"
	"github.com/g2wang/go-exercise/orgaccnt/models"
)

func Create(accountData models.AccountData) {

}

func Fetch(UUID string) models.AccountData {
	resp, err := http.Get(config.Cfg.RestURL)
	if err != nil {
		log.Fatal("error occurred, please try again")
	}
	defer resp.Body.Close()
	var accountData models.AccountData
	if err := json.NewDecoder(resp.Body).Decode(&accountData); err != nil {
		log.Fatal("error occurred, please try again")
	}
	return accountData
}

func Delete(UUID string) {

}
