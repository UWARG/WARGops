package server

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "modernc.org/sqlite"
)

type DB struct {
	*sql.DB
}

func NewDB(path string) (DB, error) {
	db, err := sql.Open("sqlite", path)
	return DB{db}, err
}

func (db DB) CreateAccount(ctx context.Context, acc NewAccount) error {
	_, err := db.DB.ExecContext(ctx, `INSERT INTO accounts (
		id,
		waterloo_id,
		name,
		source,
		allocation_date,
		expiry_date,
		active,
		creator,
		point_of_contact,
		creation_date
	) VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, $9, $10
	)`,
		acc.ID,
		acc.WaterlooID,
		acc.Name,
		acc.Source,
		acc.AllocationDate,
		acc.ExpiryDate,
		acc.Active,
		acc.Creator,
		acc.PointOfContact,
		time.Now(),
	)
	if err != nil {
		return err
	}
	return nil
}

func (db DB) ListAccounts(ctx context.Context) ([]Account, error) {
	rows, err := db.DB.QueryContext(ctx,
		`SELECT * FROM accounts
		WHERE expiry_date < DATE('now')
		UNION
		SELECT * FROM accounts
		WHERE expiry_date >= DATE('now')`)
	if err != nil {
		return nil, err
	}

	var accounts []Account

	for rows.Next() {
		var acc Account
		if err := rows.Scan(
			&acc.ID,
			&acc.WaterlooID,
			&acc.Name,
			&acc.Source,
			&acc.AllocationDate,
			&acc.ExpiryDate,
			&acc.Active,
			&acc.Creator,
			&acc.PointOfContact,
			&acc.CreationDate,
		); err != nil {
			return nil, err
		}
		accounts = append(accounts, acc)
	}

	if rows.Err() != nil {
		return nil, err
	}
	return accounts, nil
}

func (db DB) CreateTransaction(ctx context.Context, t NewTransaction) error {
	_, err := db.DB.ExecContext(ctx, `INSERT INTO transactions (
		id,
		account_id,
		creator,
		type,
		status,
		amount
	) VALUES (
		$1, $2, $3, $4, $5, $6
	)`,
		t.ID,
		t.AccountID,
		"test",
		t.Status.value,
		t.Type.value,
		t.Amount,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (db DB) ListTransactions(ctx context.Context, accountID string) ([]Transaction, error) {
	// rows, err := db.DB.QueryContext(ctx,
	// 	`SELECT * FROM transactions
	// 	WHERE account_id = $1
	// 	  AND status <> 3 -- 3 = rejected status
	// 	ORDER BY status ASC, creation_date ASC`, accountID)

	rows, err := db.DB.QueryContext(ctx,
		`SELECT * FROM transactions WHERE account_id = $1`, accountID)
	if err != nil {
		return nil, err
	}

	var transactions []Transaction

	for rows.Next() {
		var t Transaction
		if err := rows.Scan(
			&t.ID,
			&t.AccountID,
			&t.Amount,
			&t.Status,
			&t.Type,
			&t.CreationDate,
			&t.ApprovalDate,
			&t.CreationDate,
			&t.Notes,
			&t.PaymentDate,
			&t.RejectedDate,
			
		); err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}

	if rows.Err() != nil {
		return nil, err
	}

	fmt.Println("Transactions: ", transactions)
	return transactions, nil
}

func (db DB) ListAllTransactions(ctx context.Context, accountID string) ([]Transaction, error) {
	rows, err := db.DB.QueryContext(ctx,
		`SELECT * FROM transactions
		WHERE account_id = $1
		ORDER BY status ASC, creation_date ASC`, accountID)
	if err != nil {
		return nil, err
	}
	var transactions []Transaction

	for rows.Next() {
		var t Transaction
		if err := rows.Scan(
			&t.ID,
			&t.AccountID,
			&t.Amount,
			&t.Status,
			&t.Type,
			&t.CreationDate,
		); err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return transactions, nil
}

func (db DB) ListRejectedTransactions(ctx context.Context, accountID string) ([]Transaction, error) {
	rows, err := db.DB.QueryContext(ctx,
		`SELECT * FROM transactions
		WHERE account_id = $1
		  AND status = 3 -- 3 = rejected status
		ORDER BY status ASC, creation_date ASC`, accountID)
	if err != nil {
		return nil, err
	}
	var transactions []Transaction

	for rows.Next() {
		var t Transaction
		if err := rows.Scan(
			&t.ID,
			&t.AccountID,
			&t.Amount,
			&t.Status,
			&t.Type,
			&t.CreationDate,
		); err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return transactions, nil
}

func (db DB) ApproveTransaction(ctx context.Context, accountID string, transactionID string) error {
	_, err := db.DB.ExecContext(ctx, `UPDATE transactions
	SET status = 1,
		approval_date = DATE('now'),
		approved_by = $3,
		notes = $4
	WHERE id = $1
	  AND account_id = $2;`,
		transactionID,
		accountID)
	if err != nil {
		return err
	}
	return nil
}
