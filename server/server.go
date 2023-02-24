package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
)

type Server struct {
	db    DB
	bot   *discordgo.Session
	users sessions.CookieStore
}

func NewFinances(db DB, bot *discordgo.Session) *Server {
	return &Server{
		db:    db,
		bot:   bot,
		users: *sessions.NewCookieStore([]byte("changeme")),
	}
}

var _ ServerInterface = (*Server)(nil)

// Get a specific account
// (GET /accounts/{account_id}
func (f *Server) GetAccount(w http.ResponseWriter, r *http.Request, accountID string) *Response {
	account, err := f.db.GetAccount(r.Context(), accountID)
	fmt.Println("account, err", account, err)
	if err != nil {
		fmt.Println("Error: ", err)
		return &Response{
			body:        err,
			Code:        500,
			contentType: "application/json",
		}
	}

	return GetAccountJSON200Response(account)
}

// Retrieve the list of available accounts.
// (GET /accounts)
func (f *Server) ListAccounts(w http.ResponseWriter, r *http.Request) *Response {
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
func (f *Server) CreateAccount(w http.ResponseWriter, r *http.Request) *Response {
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
func (s *Server) CreateTransaction(w http.ResponseWriter, r *http.Request) *Response {
	var nt NewTransaction
	if err := json.NewDecoder(r.Body).Decode(&nt); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return &Response{
			body:        "Invalid JSON",
			Code:        400,
			contentType: "application/json",
		}
	}

	user, ok := r.Context().Value(UserKey).(goth.User)
	if !ok {
		http.Error(w, "could not find user", http.StatusBadRequest)
	}

	if err := s.db.CreateTransaction(r.Context(), nt, user.UserID); err != nil {
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
func (f *Server) ListTransactions(w http.ResponseWriter, r *http.Request, accountID string) *Response {
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
func (f *Server) ListAllTransactions(w http.ResponseWriter, r *http.Request, accountID string) *Response {
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
func (f *Server) ListRejectedTransactions(w http.ResponseWriter, r *http.Request, accountID string) *Response {
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
func (f *Server) TransactionRef(w http.ResponseWriter, r *http.Request, accountID string, transactionID string) *Response {
	return nil
}

// Approve a transaction.
// (POST /transactions/{account_id}/{transaction_id}:approve)
func (f *Server) ApproveTransaction(w http.ResponseWriter, r *http.Request, accountID string, transactionID string) *Response {
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
func (f *Server) HoldTransaction(w http.ResponseWriter, r *http.Request, accountID string, transactionID string) *Response {
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
func (f *Server) PayTransaction(w http.ResponseWriter, r *http.Request, accountID string, transactionID string) *Response {
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
func (f *Server) RejectTransaction(w http.ResponseWriter, r *http.Request, accountID string, transactionID string) *Response {
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
