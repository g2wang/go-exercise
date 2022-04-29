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
	log.Println("url:", config.Cfg.RestURL)
	resp, err := http.Get(config.Cfg.RestURL + "/v1/organisation/accounts/" + UUID)
	if err != nil {
		log.Println("error: ", err)
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
