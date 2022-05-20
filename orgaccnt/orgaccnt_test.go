package orgaccnt

import (
	"testing"

	"github.com/g2wang/go-exercise/orgaccnt/models"
	"github.com/google/uuid"
)

var createdAccountData *models.AccountData

func TestCreate(t *testing.T) {

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
	accountDataResp, err := Create(accountData)
	if err != nil {
		t.Errorf("error creating account: %v", err)
	}
	createdAccountData = accountDataResp
	if accountDataResp.ID != id {
		t.Errorf("ID mismatch after account creation. Expected: %v, Actual: %v", id, accountDataResp.ID)
	}
	if *accountDataResp.Version != 0 {
		t.Errorf("Version error after account creation. Expected: %v, Actual: %v", 0, *accountDataResp.Version)
	}
	t.Logf("account create success. id: %+v, version: %v", accountDataResp.ID, *accountDataResp.Version)
}

func TestFetch(t *testing.T) {
	if createdAccountData == nil {
		TestCreate(t)
	}
	accountData, err := Fetch(createdAccountData.ID)
	if err != nil {
		t.Errorf("Fetch error: %v", err)
	} else if accountData.ID != createdAccountData.ID {
		t.Errorf("Account Fetch error - ID mismatch. Expected: %v, Actual: %v", createdAccountData.ID, accountData.ID)
	} else if *accountData.Version != *createdAccountData.Version {
		t.Errorf("Account Fetch error - Version mismatch. Expected: %v, Actual: %v",
			*createdAccountData.Version, *accountData.Version)
	}
	t.Logf("account fetch success. id: %+v, version: %v", accountData.ID, *accountData.Version)
}

func TestDelete(t *testing.T) {
	if createdAccountData == nil {
		TestCreate(t)
	}
	success, err := Delete(createdAccountData.ID, *createdAccountData.Version)
	if err != nil {
		t.Errorf("Account Deletion error: %v", err)
	}
	if !success {
		t.Errorf("Account Deletion return value error. Expected: %v, Actual: %v", true, success)
	}
	createdAccountData = nil
}
