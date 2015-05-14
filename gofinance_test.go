package gofinance

import (
	"log"
	"testing"
)

func init() {
}

// TestMain is a sample to run an endpoint test
func TestReference(t *testing.T) {
	reference := NewReference()
	if len(reference) != 36 {
		t.Error("Expected referece with len 36")
	}
}

func TestAccountsJson(t *testing.T) {
	accs := NewAccountsBrazil()
	_, err := accs.Json()
	if err != nil {
		t.Error("Error converting Accounts to json")
	}
}

func TestAddAccount(t *testing.T) {
	accs := NewAccountsBrazil()
	// test of AddAccount
	l1 := len(accs.Asset.Childrens)
	accs.Asset.AddAccount(Account{})
	l2 := len(accs.Asset.Childrens)
	if l2 != l1+1 {
		t.Error("Error adding new Account")
	}
}

func TestAccountsSaveReload(t *testing.T) {
	accs := NewAccountsBrazil()
	accs.User = "test@gmail.com"

	err := accs.Save()
	if err != nil {
		t.Error(err)
	}

	accs2, err := GetAccountsByUser(accs.User)
	if err != nil {
		t.Error(err)
	}

	taccs, _ := (*accs).Json()
	taccs2, _ := accs2.Json()

	log.Println(string(taccs))
	log.Println(string(taccs2))
	if string(taccs) == string(taccs2) {
		t.Error("Accounts saved and accounts loaded are not equals")
	}
}
