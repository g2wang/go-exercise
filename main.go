package main

import (
	"log"

	"github.com/g2wang/go-exercise/orgaccnt"
	"github.com/g2wang/go-exercise/orgaccnt/models"
	"github.com/google/uuid"
)

func main() {
	id := testCreate()
	testFetch(id)
	testDelete(id, "0")
}

func testCreate() string {

	country := "GB"
	name := [...]string{"Guangd Wang"}

	attr := models.AccountAttributes{
		BankID:       "400301",
		BankIDCode:   "GBDSD",
		BaseCurrency: "GBP",
		Bic:          "NWBKGB23",
		Country:      &country,
		Name:         name[:],
	}

	id := uuid.NewString()
	accountData := models.AccountData{
		ID:             id,
		Type:           "accounts",
		OrganisationID: uuid.NewString(),
		Attributes:     &attr,
	}
	_, err := orgaccnt.Create(accountData)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return id
}

func testFetch(id string) {
	accountData := orgaccnt.Fetch(id)
	log.Printf("account: %+v", accountData)
	log.Printf("attributes: %+v", accountData.Attributes)
}

func testDelete(id string, version string) {
	_, err := orgaccnt.Delete(id, version)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
