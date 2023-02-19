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
	_, err := db.ExecContext(ctx, `INSERT INTO accounts (
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

// @Alexander Tsarapkine
func (db DB) GetAccount(ctx context.Context, id string) (Account, error) {
	var acc Account
	if err := db.QueryRowContext(ctx,
		`SELECT * FROM accounts
        WHERE id = $1`, id).Scan(
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
		return Account{}, err
	}

	tx, err := db.ListTransactions(ctx, id)
	if err != nil {
		return Account{}, err
	}

	for _, t := range tx {
		t = t
	}

	return acc, nil
}

func (db DB) ListAccounts(ctx context.Context) ([]Account, error) {
	rows, err := db.QueryContext(ctx,
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

func (db DB) CreateTransaction(ctx context.Context, t NewTransaction, creator string) error {
	_, err := db.ExecContext(ctx, `INSERT INTO transactions (
		id,
        account_id,
		name,
        creator,
        type,
        ref,
        status,
        amount,
        approval_date,
        approved_by,
        payment_date,
        creation_date,
        rejected_date,
        notes
	) VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14
	)`,
		t.ID,
		t.AccountID,
		t.Name,
		creator,
		t.Type.value,
		"",
		t.Status.value,
		t.Amount,
		time.Now(),
		"",
		time.Now(),
		time.Now(),
		time.Now(),
		t.Notes,
	)
	return err
}

func (db DB) ListTransactions(ctx context.Context, accountID string) ([]Transaction, error) {
	rows, err := db.QueryContext(ctx,
		`SELECT * FROM transactions
		WHERE account_id = $1
		  AND status <> 3 -- 3 = rejected status
		ORDER BY status ASC, creation_date ASC`, accountID)

	var transactions []Transaction
	var ref sql.NullString
	var creator string
	var approvedBy string

	for rows.Next() {
		var t Transaction
		if err := rows.Scan(
			&t.ID,
			&t.Name,
			&t.AccountID,
			&creator,
			&t.Type.value,
			&ref,
			&t.Status.value,
			&t.Amount,
			&t.ApprovalDate,
			&approvedBy,
			&t.PaymentDate,
			&t.CreationDate,
			&t.RejectedDate,
			&t.Notes,
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
	rows, err := db.QueryContext(ctx,
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
	rows, err := db.QueryContext(ctx,
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

func (db DB) ApproveTransaction(ctx context.Context, accountID string, transactionID string, et EditTransaction) error {
	_, err := db.ExecContext(ctx, `UPDATE transactions
	SET status = 1,
		approval_date = DATE('now'),
		approved_by = $3,
		notes = $4
	WHERE id = $1
	  AND account_id = $2`,
		transactionID,
		accountID,
		et.Approver,
		et.Notes)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (db DB) RejectTransaction(ctx context.Context, accountID string, transactionID string, et EditTransaction) error {
	_, err := db.ExecContext(ctx, `UPDATE transactions
	SET status = 3,
		rejected_date = DATE('now'),
		approved_by = $3,
		notes = $4
		
	WHERE id = $1
	  AND account_id = $2`,
		transactionID,
		accountID,
		et.Approver,
		et.Notes)
	if err != nil {
		return err
	}
	return nil
}

func (db DB) PayTransaction(ctx context.Context, accountID string, transactionID string, et EditTransaction) error {
	_, err := db.ExecContext(ctx, `UPDATE transactions
	SET status = 2,
		payment_date = DATE('now'),
		ref = $5,
		notes = $4,
		approved_by = $3
	WHERE id = $1
	  AND account_id = $2`,
		transactionID,
		accountID,
		et.Approver,
		et.Notes,
		"ref")
	if err != nil {
		return err
	}
	return nil
}

func (db DB) HoldTransaction(ctx context.Context, accountID string, transactionID string, et EditTransaction) error {
	_, err := db.ExecContext(ctx, `UPDATE transactions
SET status = 0,
	approved_by = $3,
    notes = $4
WHERE id = $1
  AND account_id = $2`,
		transactionID,
		accountID,
		et.Approver,
		et.Notes)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
