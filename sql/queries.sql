--- create an account
INSERT INTO accounts (
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
);

-- create a transaction
INSERT INTO transactions (
  id,
  account_id,
  amount,
  type,
  creation_date,
  creator,
  status,
  notes
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
);


--- list all accounts, list all active accounts (oldest first), then all expired accounts (most recently expired first)
SELECT * FROM accounts
WHERE expiry_date > DATE('now')
ORDER BY allocation_date ASC
UNION
SELECT * FROM accounts
WHERE expiry_date <= DATE('now')
ORDER BY expiry_date DESC;

-- list all transactions for an account, exclude rejected
-- puts all transactions that need a decision first (they are pending decisions from leads, but don't impact accounts) (type is created)
-- puts all transactions that are pending next (approved, but not finalized)
-- puts all other transactions next, with oldest (which are not rejected) first
SELECT * FROM transactions
WHERE account_id = $1
  AND status <> 3 -- 3 = rejected status
ORDER BY status ASC, creation_date ASC;

-- list all approved and paid transactions (ignores transactions waiting for a decision and those rejected)
-- puts all transactions that are pending first (approved, but not finalized)
-- puts all other transactions next, with oldest (which are not rejected) first
SELECT * FROM transactions
WHERE account_id = $1
  AND status <> 3 -- 3 = rejected status
  AND status <> 0 -- 0 = created status
ORDER BY status ASC, creation_date ASC;

-- list all rejected transactions
SELECT * FROM transactions
WHERE account_id = $1
  AND status = 3
ORDER BY creation_date ASC;

-- approve transaction
UPDATE transactions
SET status = 1,
    approval_date = DATE('now'),
    approved_by = $3,
    notes = $4
WHERE id = $1
  AND account_id = $2;

-- reject transaction
UPDATE transactions
SET status = 3,
    rejected_date = DATE('now'),
    notes = $3
WHERE id = $1
  AND account_id = $2;

-- mark transaction as paid
UPDATE transactions
SET status = 2,
    payment_date = DATE('now'),
    ref = $3,
    notes = $4
WHERE id = $1
  AND account_id = $2;

-- hold transaction
UPDATE transactions
SET status = 0,
    notes = $3
WHERE id = $1
  AND account_id = $2;


-- Allocated Funds
SELECT SUM(amount) AS total_amount
FROM transactions
WHERE account_id = $1
AND type = 0
AND status = 2;

-- Account Balance
SELECT SUM(amount) AS total_amount
FROM transactions
WHERE account_id = $1
AND status = 2;
