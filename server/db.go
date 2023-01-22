package server

import (
	"context"
	"database/sql"
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
		source,
		allocation_date,
		expiry_date,
		active,
		creator,
		point_of_contact,
		creation_date
	) VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, $9
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
		amount,
		status,
		type,
	) VALUES (
		$1, $2, $3, $4, $5,
	)`,
		t.ID,
		t.AccountID,
		t.Amount,
		t.Status,
		t.Type,
		time.Now(),
	)
	if err != nil {
		return err
	}
	return nil
}

func (db DB) ListTransactions(ctx context.Context, accountID string) ([]Transaction, error) {
	rows, err := db.DB.QueryContext(ctx,
		`SELECT * FROM transactions
		WHERE account_id = $1
		  AND status <> 3 -- 3 = rejected status
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

func (db Db) ListRejectedTransactions(ctx context.Context, accountID string) ([]Transaction, error) {
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

