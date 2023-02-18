package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Finances struct {
	db DB
}

func NewFinances(db DB) *Finances {
	return &Finances{
		db: db,
	}
}

var _ ServerInterface = (*Finances)(nil)

// Retrieve the list of available accounts.
// (GET /accounts)
func (f *Finances) ListAccounts(w http.ResponseWriter, r *http.Request) *Response {
	accounts, err := f.db.ListAccounts(r.Context())
	if err != nil {
		fmt.Println("Error: ", err)
		return &Response{
			body:        err,
			Code:        500,
			contentType: "application/json",
		}
	}

	return ListAccountsJSON200Response(accounts)
}

// Create a new account.
// (POST /accounts)
func (f *Finances) CreateAccount(w http.ResponseWriter, r *http.Request) *Response {
	newAccount := NewAccount{}
	if err := json.NewDecoder(r.Body).Decode(&newAccount); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return &Response{
			body:        "Invalid JSON",
			Code:        400,
			contentType: "application/json",
		}
	}
	if err := f.db.CreateAccount(r.Context(), newAccount); err != nil {
		fmt.Println("Error: ", err)
		return &Response{
			body:        err,
			Code:        500,
			contentType: "application/json",
		}
	}

	return nil
}

// Create a new transaction.
// (POST /transactions)
func (f *Finances) CreateTransaction(w http.ResponseWriter, r *http.Request) *Response {
	var nt NewTransaction
	if err := json.NewDecoder(r.Body).Decode(&nt); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return &Response{
			body:        "Invalid JSON",
			Code:        400,
			contentType: "application/json",
		}
	}
	if err := f.db.CreateTransaction(r.Context(), nt, "test"); err != nil {
		fmt.Println("Error: ", err)
		return &Response{
			body:        err,
			Code:        500,
			contentType: "application/json",
		}
	}
	return nil
}

// Retrieve the active transactions for an account.
// (GET /transactions/{account_id})
func (f *Finances) ListTransactions(w http.ResponseWriter, r *http.Request, accountID string) *Response {
	transactions, err := f.db.ListTransactions(r.Context(), accountID)
	if err != nil {
		fmt.Println("Error: ", err)
		return &Response{
			body:        err,
			Code:        500,
			contentType: "application/json",
		}
	}
	return ListTransactionsJSON200Response(transactions)
}

// Retrieve all transactions, including those that are currently a request and not approved.
// (GET /transactions/{account_id}/all)
func (f *Finances) ListAllTransactions(w http.ResponseWriter, r *http.Request, accountID string) *Response {
	transactions, err := f.db.ListAllTransactions(r.Context(), accountID)
	if err != nil {
		return &Response{
			body:        err,
			Code:        500,
			contentType: "application/json",
		}
	}

	return ListAllTransactionsJSON200Response(transactions)
}

// Retrieve all rejected transactions for an account.
// (GET /transactions/{account_id}/rejected)
func (f *Finances) ListRejectedTransactions(w http.ResponseWriter, r *http.Request, accountID string) *Response {
	transactions, err := f.db.ListRejectedTransactions(r.Context(), accountID)
	if err != nil {
		return &Response{
			body:        err,
			Code:        500,
			contentType: "application/json",
		}
	}

	return ListRejectedTransactionsJSON200Response(transactions)
}

// Get transaction documents
// (GET /transactions/{account_id}/{transaction_id}/ref)
func (f *Finances) TransactionRef(w http.ResponseWriter, r *http.Request, accountID string, transactionID string) *Response {
	return nil
}

// Approve a transaction.
// (POST /transactions/{account_id}/{transaction_id}:approve)
func (f *Finances) ApproveTransaction(w http.ResponseWriter, r *http.Request, accountID string, transactionID string) *Response {
	editTransaction := EditTransaction{}
	if err := json.NewDecoder(r.Body).Decode(&editTransaction); err != nil {
		fmt.Println("Error: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return &Response{
			body:        "Invalid JSON",
			Code:        400,
			contentType: "application/json",
		}
	}
	fmt.Println("Notes: ", editTransaction.Notes)
	err := f.db.ApproveTransaction(r.Context(), accountID, transactionID, editTransaction)
	if err != nil {
		fmt.Println("Error: ", err)
		return &Response{
			body:        err,
			Code:        500,
			contentType: "application/json",
		}
	}

	return nil
}

// Hold back a transaction, and reset it to pending
// (POST /transactions/{account_id}/{transaction_id}:hold)
func (f *Finances) HoldTransaction(w http.ResponseWriter, r *http.Request, accountID string, transactionID string) *Response {
	editTransaction := EditTransaction{}
	if err := json.NewDecoder(r.Body).Decode(&editTransaction); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return &Response{
			body:        err,
			Code:        400,
			contentType: "application/json",
		}
	}
	fmt.Println("Notes: ", editTransaction.Notes)

	err := f.db.HoldTransaction(r.Context(), accountID, transactionID, editTransaction)
	if err != nil {
		return &Response{
			body:        err,
			Code:        500,
			contentType: "application/json",
		}
	}
	return nil
}

// Mark a transaction as paid.
// (POST /transactions/{account_id}/{transaction_id}:pay)
func (f *Finances) PayTransaction(w http.ResponseWriter, r *http.Request, accountID string, transactionID string) *Response {
	editTransaction := EditTransaction{}
	if err := json.NewDecoder(r.Body).Decode(&editTransaction); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return &Response{
			body:        "Invalid JSON",
			Code:        400,
			contentType: "application/json",
		}
	}
	fmt.Println("Notes: ", editTransaction.Notes)

	err := f.db.PayTransaction(r.Context(), accountID, transactionID, editTransaction)
	if err != nil {
		return &Response{
			body:        err,
			Code:        500,
			contentType: "application/json",
		}
	}
	return nil
}

// Reject a transaction.
// (POST /transactions/{account_id}/{transaction_id}:reject)
func (f *Finances) RejectTransaction(w http.ResponseWriter, r *http.Request, accountID string, transactionID string) *Response {
	editTransaction := EditTransaction{}
	if err := json.NewDecoder(r.Body).Decode(&editTransaction); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return &Response{
			body:        "Invalid JSON",
			Code:        400,
			contentType: "application/json",
		}
	}
	fmt.Println("Notes: ", editTransaction.Notes)

	err := f.db.RejectTransaction(r.Context(), accountID, transactionID, editTransaction)
	if err != nil {
		fmt.Println("Error: ", err)
		return &Response{
			body:        err,
			Code:        500,
			contentType: "application/json",
		}
	}
	return nil
}
