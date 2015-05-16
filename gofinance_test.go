package gofinance

import (
	"log"
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
	log.Println(accs)

	refAtivo, err := accs.GetAccountRefByName("Ativo")
	log.Println("Ativo", refAtivo)
	if err == nil {
		if refAtivo != accs.Asset.Reference {
			t.Error("Account with name Ativo Found Fail.")
		}
	} else {
		t.Error("Account with name Ativo not Found.")
	}

	refCaixa, err := accs.GetAccountRefByName("Caixa")
	log.Println("Caixa", refCaixa)
	if err == nil {
		if refCaixa != accs.Asset.Childrens[0].Reference {
			t.Error("Account with name Caixa Found Fail.")
		}
	} else {
		t.Error("Account with name Caixa not Found.")
	}

	refPrestacoes, err := accs.GetAccountRefByName("Prestacoes")
	log.Println("Prestacoes", refPrestacoes)
	if err == nil {
		if refPrestacoes != accs.Liability.Childrens[2].Reference {
			t.Error("Account with name Prestacoes Found Fail.")
		}
	} else {
		t.Error("Account with name Prestacoes not Found.")
	}

	refSalario, err := accs.GetAccountRefByName("Salario")
	log.Println("Salario", refSalario)
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
