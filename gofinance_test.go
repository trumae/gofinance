package gofinance

import (
	"github.com/shopspring/decimal"
	"testing"
)

func init() {
}

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

	accs2 := Accounts{}
	accs2, err = GetAccountsByUser(accs.User)
	if err != nil {
		t.Error(err)
	}

	if string(accs.User) != string(accs2.User) {
		t.Error("Accounts saved and accounts loaded are not equals")
	}
}

func TestHasChildrens(t *testing.T) {
	accs := NewAccountsBrazil()

	if !accs.Asset.HasChildrens() {
		t.Error("Account Ativos has childrens")
	}

	if accs.Asset.Childrens[0].HasChildrens() {
		t.Error("Account Caixa not has childrens")
	}
}

func TestGetAccountRefByName(t *testing.T) {
	//func (accs *Accounts) GetAccountRefByName(name string) (string, error)
	accs := NewAccountsBrazil()

	refAtivo, err := accs.GetAccountRefByName("Ativo")
	if err == nil {
		if refAtivo != accs.Asset.Reference {
			t.Error("Account with name Ativo Found Fail.")
		}
	} else {
		t.Error("Account with name Ativo not Found.")
	}

	refCaixa, err := accs.GetAccountRefByName("Caixa")
	if err == nil {
		if refCaixa != accs.Asset.Childrens[0].Reference {
			t.Error("Account with name Caixa Found Fail.")
		}
	} else {
		t.Error("Account with name Caixa not Found.")
	}

	refPrestacoes, err := accs.GetAccountRefByName("Prestacoes")
	if err == nil {
		if refPrestacoes != accs.Liability.Childrens[2].Reference {
			t.Error("Account with name Prestacoes Found Fail.")
		}
	} else {
		t.Error("Account with name Prestacoes not Found.")
	}

	refSalario, err := accs.GetAccountRefByName("Salario")
	if err == nil {
		if refSalario != accs.Income.Childrens[0].Reference {
			t.Error("Account with name Salario Found Fail.")
		}
	} else {
		t.Error("Account with name Salario not Found.")
	}

	_, err = accs.GetAccountRefByName("XPTO")
	if err == nil {
		t.Error("Account with name XPTO found, but not exist.")
	}
}

func TestAccountByRef(t *testing.T) {
	accs := NewAccountsBrazil()

	refAtivo := accs.Asset.Reference
	refCaixa := accs.Asset.Childrens[0].Reference

	accAtivo, err := accs.GetAccountByRef(refAtivo)
	if err != nil {
		t.Error(err)
	}
	if accAtivo.Reference != refAtivo {
		t.Error("Error in Account returned")
	}

	accCaixa, err := accs.GetAccountByRef(refCaixa)
	if err != nil {
		t.Error(err)
	}
	if accCaixa.Reference != refCaixa {
		t.Error("Error in Account returned")
	}

	_, err = accs.GetAccountByRef("ref")
	if err == nil {
		t.Error("Found not exist ref")
	}
}

func TestNewEntryAndValid(t *testing.T) {
	accs := NewAccountsBrazil()
	accs.User = "test2@gmail.com"
	accs.Save()

	refAtivo := accs.Asset.Childrens[0].Reference

	ent := NewEntry(accs, "Test of entry")
	if !ent.Valid() {
		t.Error("Error in Valid test")
	}

	ent.AddDebit(refAtivo, decimal.NewFromFloat(10.0), "first debit")
	if ent.Valid() {
		t.Error("Error in Valid test")
	}

	ent.AddDebit(refAtivo, decimal.NewFromFloat(5.0), "second debit")
	if ent.Valid() {
		t.Error("Error in Valid test")
	}

	ent.AddCredit(refAtivo, decimal.NewFromFloat(7.0), "first credit")
	if ent.Valid() {
		t.Error("Error in Valid test")
	}

	ent.AddCredit(refAtivo, decimal.NewFromFloat(8.0), "second credit")
	if !ent.Valid() {
		t.Error("Error in Valid test")
	}
}

func TestEntrySave(t *testing.T) {
	accs := NewAccountsBrazil()
	accs.User = "test3@gmail.com"
	accs.Save()

	refCaixa, _ := accs.GetAccountRefByName("Caixa")

	ent := NewEntry(accs, "Save Test")
	ent.AddDebit(refCaixa, decimal.NewFromFloat(10.0), "first debit")
	ent.AddDebit(refCaixa, decimal.NewFromFloat(5.0), "second debit")
	ent.AddCredit(refCaixa, decimal.NewFromFloat(7.0), "first credit")
	ent.AddCredit(refCaixa, decimal.NewFromFloat(8.0), "second credit")

	err := ent.Save()
	if err != nil {
		t.Error(err)
	}

	ent2 := &Entry{}
	ent2, err = GetEntryByRef(ent.Reference)
	if err != nil {
		t.Error(err)
	}

	if string(ent.Reference) != string(ent2.Reference) {
		t.Error("Entries saved not are equals")
	}
}

func TestDeleteEntry(t *testing.T) {
	accs := NewAccountsBrazil()
	accs.User = "test4@gmail.com"
	accs.Save()

	ent := NewEntry(accs, "entry1")
	ent.Save()

	ent2, err := GetEntryByRef(ent.Reference)
	if err != nil {
		t.Error("Error saving", err)
	}

	if ent.Reference != ent2.Reference {
		t.Error(err)
	}
}

////////// Benchmarks ////////////////////////

func BenchmarkAccountsSave(t *testing.B) {
	accs := NewAccountsBrazil()
	accs.User = "test@gmail.com"
	accs.Save()
}

func BenchmarkEntrySave(t *testing.B) {
	accs := NewAccountsBrazil()
	accs.User = "test@gmail.com"
	accs.Save()

	refCaixa, _ := accs.GetAccountRefByName("Caixa")
	refSalario, _ := accs.GetAccountRefByName("Salario")

	ent := NewEntry(accs, "Save Test")
	ent.AddDebit(refCaixa, decimal.NewFromFloat(10.0), "Salario")
	ent.AddCredit(refSalario, decimal.NewFromFloat(10.0), "Salario")

	ent.Save()
}
