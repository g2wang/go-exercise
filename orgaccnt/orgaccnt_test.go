package orgaccnt

import (
	"testing"

	"github.com/g2wang/go-exercise/orgaccnt/models"
	"github.com/google/go-cmp/cmp"
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
	version := int64(0)
	accountData := models.AccountData{
		ID:             id,
		Type:           "accounts",
		OrganisationID: uuid.NewString(),
		Version:        &version,
		Attributes:     &attr,
	}
	accountDataResp, err := Create(accountData)
	if err != nil {
		t.Errorf("error creating account: %v", err)
	}
	createdAccountData = accountDataResp
	if d := cmp.Diff(*createdAccountData, accountData); d != "" {
		t.Errorf("unexpected difference in featched acccount data:\n%v", d)
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
	} else if d := cmp.Diff(*createdAccountData, *accountData); d != "" {
		t.Errorf("unexpected difference in featched acccount data:\n%v", d)
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
	_, err = Fetch(createdAccountData.ID)
	if err == nil {
		t.Errorf("Deleted but still fetched error: %v", err)
	}
	createdAccountData = nil
}
