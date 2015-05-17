package gofinance

import (
	"code.google.com/p/go-uuid/uuid"
	"encoding/json"
	"errors"
	"github.com/boltdb/bolt"
	"github.com/shopspring/decimal"
	"log"
	"time"
)

const (
	Debit = iota
	Credit

	AccountsTypeBrazil = iota
	AccountTypeUSA
)

var (
	// errors
	notFound = errors.New("NotFound")
	notValid = errors.New("NotValid")

	dbFilename = "accounts.db"
	bAccounts  = "accounts"
	bEntries   = "entries"
)

func init() {
	db, err := bolt.Open(dbFilename, 0600, nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bAccounts))
		if err != nil {
			log.Fatalln("create bucket:", bAccounts, err)
		}
		_, err = tx.CreateBucketIfNotExists([]byte(bEntries))
		if err != nil {
			log.Fatalln("create bucket:", bEntries, err)
		}
		return nil
	})
}

type Account struct {
	Reference string
	Name      string
	Type      int
	Childrens []Account
	Info      string
	CreatedAt time.Time
	Balance   decimal.Decimal
}

type Accounts struct {
	User      string
	Type      int /// AccountType
	Asset     Account
	Liability Account
	Income    Account
	Expense   Account
}

type UnitEntry struct {
	Value      decimal.Decimal
	RefAccount string
	Info       string
}

type Entry struct {
	User      string
	Reference string
	Debits    []UnitEntry
	Credits   []UnitEntry
	CreatedAt time.Time
	Info      string
}

func NewReference() string {
	return uuid.NewRandom().String()
}

func (acc *Account) AddAccount(child Account) {
	acc.Childrens = append(acc.Childrens, child)
}

func NewAccountsBrazil() *Accounts {
	//ATIVO
	ativo := Account{
		Reference: NewReference(),
		Name:      "Ativo",
		Type:      Debit,
		Info:      "",
		CreatedAt: time.Now()}

	ativo.AddAccount(Account{
		Reference: NewReference(),
		Name:      "Caixa",
		Type:      Debit,
		Info:      "",
		CreatedAt: time.Now()})

	ativo.AddAccount(Account{
		Reference: NewReference(),
		Name:      "Bancos",
		Type:      Debit,
		Info:      "",
		CreatedAt: time.Now()})
	ativo.AddAccount(Account{
		Reference: NewReference(),
		Name:      "Receber",
		Type:      Debit,
		Info:      "",
		CreatedAt: time.Now()})
	ativo.AddAccount(Account{
		Reference: NewReference(),
		Name:      "Bens",
		Type:      Debit,
		Info:      "",
		CreatedAt: time.Now()})

	//PASSIVO
	passivo := Account{
		Reference: NewReference(),
		Name:      "Passivo",
		Type:      Credit,
		Info:      "",
		CreatedAt: time.Now()}
	passivo.AddAccount(Account{
		Reference: NewReference(),
		Name:      "Pagar",
		Type:      Credit,
		Info:      "",
		CreatedAt: time.Now()})
	passivo.AddAccount(Account{
		Reference: NewReference(),
		Name:      "Cartao de Credito",
		Type:      Credit,
		Info:      "",
		CreatedAt: time.Now()})
	passivo.AddAccount(Account{
		Reference: NewReference(),
		Name:      "Prestacoes",
		Type:      Credit,
		Info:      "",
		CreatedAt: time.Now()})

	//RECEITAS
	receitas := Account{
		Reference: NewReference(),
		Name:      "Receitas",
		Type:      Credit,
		Info:      "",
		CreatedAt: time.Now()}
	receitas.AddAccount(Account{
		Reference: NewReference(),
		Name:      "Salario",
		Type:      Credit,
		Info:      "",
		CreatedAt: time.Now()})
	receitas.AddAccount(Account{
		Reference: NewReference(),
		Name:      "Rendimentos",
		Type:      Credit,
		Info:      "",
		CreatedAt: time.Now()})
	receitas.AddAccount(Account{
		Reference: NewReference(),
		Name:      "Outras Receitas",
		Type:      Credit,
		Info:      "",
		CreatedAt: time.Now()})

	// DESPESAS
	despesas := Account{
		Reference: NewReference(),
		Name:      "Despesas",
		Type:      Debit,
		Info:      "",
		CreatedAt: time.Now()}
	despesas.AddAccount(Account{
		Reference: NewReference(),
		Name:      "Fiscais",
		Type:      Debit,
		Info:      "",
		CreatedAt: time.Now()})
	despesas.AddAccount(Account{
		Reference: NewReference(),
		Name:      "Aluguel",
		Type:      Debit,
		Info:      "",
		CreatedAt: time.Now()})
	despesas.AddAccount(Account{
		Reference: NewReference(),
		Name:      "Agua",
		Type:      Debit,
		Info:      "",
		CreatedAt: time.Now()})
	despesas.AddAccount(Account{
		Reference: NewReference(),
		Name:      "Luz",
		Type:      Debit,
		Info:      "",
		CreatedAt: time.Now()})
	despesas.AddAccount(Account{
		Reference: NewReference(),
		Name:      "Telefone",
		Type:      Debit,
		Info:      "",
		CreatedAt: time.Now()})
	despesas.AddAccount(Account{
		Reference: NewReference(),
		Name:      "Internet",
		Type:      Debit,
		Info:      "",
		CreatedAt: time.Now()})
	despesas.AddAccount(Account{
		Reference: NewReference(),
		Name:      "Cartao de Credito",
		Type:      Debit,
		Info:      "",
		CreatedAt: time.Now()})
	despesas.AddAccount(Account{
		Reference: NewReference(),
		Name:      "Salarios",
		Type:      Debit,
		Info:      "",
		CreatedAt: time.Now()})
	despesas.AddAccount(Account{
		Reference: NewReference(),
		Name:      "Frete",
		Type:      Debit,
		Info:      "",
		CreatedAt: time.Now()})

	return &Accounts{Asset: ativo,
		Liability: passivo,
		Income:    receitas,
		Expense:   despesas,
		Type:      AccountsTypeBrazil}
}

func (acc *Account) getAccountRefByNameRec(name string) (string, error) {
	if acc.Name == name {
		return acc.Reference, nil
	}
	for _, val := range acc.Childrens {
		ref, err := val.getAccountRefByNameRec(name)
		if err == nil {
			return ref, nil
		}
	}
	return "", notFound
}

func (acc *Account) getAccountByRefRec(ref string) (*Account, error) {
	if acc.Reference == ref {
		return acc, nil
	}
	for _, val := range acc.Childrens {
		nacc, err := val.getAccountByRefRec(ref)
		if err == nil {
			return nacc, nil
		}
	}
	return nil, notFound
}

func (accs *Accounts) GetAccountRefByName(name string) (string, error) {
	val, err := accs.Asset.getAccountRefByNameRec(name)
	if err == nil {
		return val, err
	}
	val, err = accs.Liability.getAccountRefByNameRec(name)
	if err == nil {
		return val, err
	}
	val, err = accs.Income.getAccountRefByNameRec(name)
	if err == nil {
		return val, err
	}
	val, err = accs.Expense.getAccountRefByNameRec(name)
	if err == nil {
		return val, err
	}

	return "", notFound
}

func (accs *Accounts) GetAccountByRef(ref string) (*Account, error) {
	val, err := accs.Asset.getAccountByRefRec(ref)
	if err == nil {
		return val, err
	}

	val, err = accs.Liability.getAccountByRefRec(ref)
	if err == nil {
		return val, err
	}

	val, err = accs.Income.getAccountByRefRec(ref)
	if err == nil {
		return val, err
	}

	val, err = accs.Expense.getAccountByRefRec(ref)
	if err == nil {
		return val, err
	}

	return &Account{}, notFound
}

func (acc *Account) HasChildrens() bool {
	return len(acc.Childrens) != 0
}

func NewEntry(accs *Accounts, info string) *Entry {
	return &Entry{
		User:      accs.User,
		Reference: NewReference(),
		Info:      info,
		CreatedAt: time.Now()}
}

func (ent *Entry) AddDebit(
	refacc string,
	value decimal.Decimal,
	info string) {
	ue := UnitEntry{
		RefAccount: refacc,
		Value:      value,
		Info:       info}
	ent.Debits = append(ent.Debits, ue)
}

func (ent *Entry) AddCredit(
	refacc string,
	value decimal.Decimal,
	info string) {
	ue := UnitEntry{
		RefAccount: refacc,
		Value:      value,
		Info:       info}
	ent.Credits = append(ent.Credits, ue)
}

func (ent *Entry) Valid() bool {
	accs, err := GetAccountsByUser(ent.User)
	if err != nil {
		return false
	}

	if len(ent.Reference) == 0 {
		return false
	}

	sumDebit := decimal.NewFromFloat(0.0)
	for _, uent := range ent.Debits {
		sumDebit = sumDebit.Add(uent.Value)
		_, err := accs.GetAccountByRef(uent.RefAccount)
		if err != nil {
			return false
		}
	}

	sumCredit := decimal.NewFromFloat(0.0)
	for _, uent := range ent.Credits {
		sumCredit = sumCredit.Add(uent.Value)
		_, err := accs.GetAccountByRef(uent.RefAccount)
		if err != nil {
			return false
		}
	}

	if !sumCredit.Equals(sumDebit) {
		return false
	}

	return true
}

// Persistence
func GetAccountsByUser(name string) (Accounts, error) {
	db, err := bolt.Open(dbFilename, 0600, nil)
	if err != nil {
		return Accounts{}, err
	}
	defer db.Close()
	accs := Accounts{}

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bAccounts))
		raccs := b.Get([]byte(name))
		err := json.Unmarshal(raccs, &accs)
		return err
	})

	return accs, err
}

func (accs *Accounts) Save() error {
	db, err := bolt.Open(dbFilename, 0600, nil)
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bAccounts))
		json, err := accs.Json()
		if err != nil {
			log.Println(err)
		}
		err = b.Put([]byte(accs.User), json)
		return err
	})

	return err
}

func (accs *Accounts) Remove() error {
	//TODO
	return errors.New("Remove Accounts not implemented")
}

func GetEntryByRef(ref string) (*Entry, error) {
	db, err := bolt.Open(dbFilename, 0600, nil)
	if err != nil {
		return &Entry{}, err
	}
	defer db.Close()
	entry := Entry{}

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bEntries))
		ent := b.Get([]byte(ref))
		err := json.Unmarshal(ent, &entry)
		return err
	})

	return &entry, err
}
func (ent *Entry) Save() error {
	if !ent.Valid() {
		return notValid
	}

	db, err := bolt.Open(dbFilename, 0600, nil)
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bEntries))
		json, err := ent.Json()
		if err != nil {
			log.Println(err)
		}
		err = b.Put([]byte(ent.Reference), json)
		return err
	})

	return err
}

func (ent *Entry) Remove(ref string) error {
	//TODO
	return nil
}

// Marshal
func (accs *Accounts) Json() ([]byte, error) {
	b, err := json.Marshal(accs)
	if err != nil {
		log.Println("error:", err)
		return nil, err
	}
	return b, nil
}

func (acc *Account) Json() ([]byte, error) {
	b, err := json.Marshal(acc)
	if err != nil {
		log.Println("error:", err)
		return nil, err
	}
	return b, nil
}

func (ent *Entry) Json() ([]byte, error) {
	b, err := json.Marshal(ent)
	if err != nil {
		log.Println("error:", err)
		return nil, err
	}
	return b, nil
}
