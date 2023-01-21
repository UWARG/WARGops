package server

import (
	"fmt"
	"net/http"
)

type Finances struct {
	Accounts     map[string]Account
	Transactions map[string]Transaction
}

func NewFinances() *Finances {
	return &Finances{
		Accounts:     make(map[string]Account),
		Transactions: make(map[string]Transaction),
	}
}

var _ ServerInterface = (*Finances)(nil)

// Retrieve the list of available accounts.
// (GET /accounts)
func (f *Finances) ListAccounts(w http.ResponseWriter, r *http.Request) *Response {
	fmt.Println("called!")
	return ListAccountsJSON200Response([]Account{
		{
			PointOfContact: "h2ali@uwaterloo.ca",
		},
	})
}

// Create a new account.
// (POST /accounts)
func (f *Finances) CreateAccount(w http.ResponseWriter, r *http.Request) *Response { return nil }

// Create a new transaction.
// (POST /transactions)
func (f *Finances) CreateTransaction(w http.ResponseWriter, r *http.Request) *Response { return nil }

// Retrieve the active transactions for an account.
// (GET /transactions/{account_id})
func (f *Finances) ListTransactions(w http.ResponseWriter, r *http.Request, accountID string) *Response {
	return nil
}

// Retrieve all transactions, including those that are currently a request and not approved.
// (GET /transactions/{account_id}/all)
func (f *Finances) ListAllTransactions(w http.ResponseWriter, r *http.Request, accountID string) *Response {
	return nil
}

// Retrieve all rejected transactions for an account.
// (GET /transactions/{account_id}/rejected)
func (f *Finances) ListRejectedTransactions(w http.ResponseWriter, r *http.Request, accountID string) *Response {
	return nil
}

// Get transaction documents
// (GET /transactions/{account_id}/{transaction_id}/ref)
func (f *Finances) TransactionRef(w http.ResponseWriter, r *http.Request, accountID string, transactionID string) *Response {
	return nil
}

// Approve a transaction.
// (POST /transactions/{account_id}/{transaction_id}:approve)
func (f *Finances) ApproveTransaction(w http.ResponseWriter, r *http.Request, accountID string, transactionID string) *Response {
	return nil
}

// Hold back a transaction, and reset it to pending
// (POST /transactions/{account_id}/{transaction_id}:hold)
func (f *Finances) HoldTransaction(w http.ResponseWriter, r *http.Request, accountID string, transactionID string) *Response {
	return nil
}

// Mark a transaction as paid.
// (POST /transactions/{account_id}/{transaction_id}:pay)
func (f *Finances) PayTransaction(w http.ResponseWriter, r *http.Request, accountID string, transactionID string) *Response {
	return nil
}

// Reject a transaction.
// (POST /transactions/{account_id}/{transaction_id}:reject)
func (f *Finances) RejectTransaction(w http.ResponseWriter, r *http.Request, accountID string, transactionID string) *Response {
	return nil
}
