package gofinance

import (
	"code.google.com/p/go-uuid/uuid"
	"encoding/json"
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
	Reference string
	Value     decimal.Decimal
	Info      string
}

type Entry struct {
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
	receitas.AddAccount(Account{
		Reference: NewReference(),
		Name:      "Fiscais",
		Type:      Debit,
		Info:      "",
		CreatedAt: time.Now()})
	receitas.AddAccount(Account{
		Reference: NewReference(),
		Name:      "Aluguel",
		Type:      Debit,
		Info:      "",
		CreatedAt: time.Now()})
	receitas.AddAccount(Account{
		Reference: NewReference(),
		Name:      "Agua",
		Type:      Debit,
		Info:      "",
		CreatedAt: time.Now()})
	receitas.AddAccount(Account{
		Reference: NewReference(),
		Name:      "Luz",
		Type:      Debit,
		Info:      "",
		CreatedAt: time.Now()})
	receitas.AddAccount(Account{
		Reference: NewReference(),
		Name:      "Telefone",
		Type:      Debit,
		Info:      "",
		CreatedAt: time.Now()})
	receitas.AddAccount(Account{
		Reference: NewReference(),
		Name:      "Internet",
		Type:      Debit,
		Info:      "",
		CreatedAt: time.Now()})
	receitas.AddAccount(Account{
		Reference: NewReference(),
		Name:      "Cartao de Credito",
		Type:      Debit,
		Info:      "",
		CreatedAt: time.Now()})
	receitas.AddAccount(Account{
		Reference: NewReference(),
		Name:      "Salarios",
		Type:      Debit,
		Info:      "",
		CreatedAt: time.Now()})
	receitas.AddAccount(Account{
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

func GetAccountsByName(name string) Accounts {
	// TODO
	return Accounts{}
}

func (accs *Accounts) GetAccountRefByName(name string) string {
	// TODO
	return ""
}

func (acc *Account) HasChildrens() bool {
	// TODO
	return false
}

func (accs *Accounts) HasAccountByRef() bool {
	//TODO
	return false
}

func (accs *Accounts) GetAccountByRef() (Account, error) {
	//TODO
	return Account{}, nil
}

func NewEntry(info string) Entry {
	//TODO
	return Entry{Info: info,
		CreatedAt: time.Now()}
}

func (ent *Entry) AddDebit(ref string,
	value decimal.Decimal,
	info string) {
	//TODO
	ue := UnitEntry{Reference: ref,
		Value: value,
		Info:  info}
	ent.Debits = append(ent.Debits, ue)
}

func (ent *Entry) AddCredit(ref string,
	value decimal.Decimal,
	info string) {
	//TODO
	ue := UnitEntry{Reference: ref,
		Value: value,
		Info:  info}
	ent.Credits = append(ent.Credits, ue)
}

func (ent Entry) Valid() bool {
	//TODO
	return false
}

// Persistence

func (accs *Accounts) Save() error {
	// TODO
	return nil
}

func (accs *Accounts) Remove() error {
	//TODO
	return nil
}

func (ent *Entry) Save() error {
	// TODO
	return nil
}

func (ent *Entry) Remove() error {
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
