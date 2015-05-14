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

func TestAddAccount(t *testing.T) {
	accs := NewAccountsBrazil()
	m, err := accs.Json()
	if err != nil {
		t.Error("Error converting Accounts to json")
	}
	log.Println(string(m))

	// test of AddAccount
	l1 := len(accs.Asset.Childrens)
	accs.Asset.AddAccount(Account{})
	l2 := len(accs.Asset.Childrens)
	if l2 != l1+1 {
		t.Error("Error adding new Account")
	}
}
