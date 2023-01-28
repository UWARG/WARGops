CREATE TABLE accounts (
	id VARCHAR(64) NOT NULL PRIMARY KEY,
	waterloo_id VARCHAR(255) NOT NULL UNIQUE,
	name VARCHAR(255) NOT NULL UNIQUE,
	source VARCHAR(255) NOT NULL,
	allocation_date TIMESTAMP NOT NULL,
	expiry_date TIMESTAMP NOT NULL,
	active BOOLEAN NOT NULL,
	creator VARCHAR(32) NOT NULL,
	point_of_contact VARCHAR(64) NOT NULL,
	creation_date TIMESTAMP NOT NULL
);

CREATE TABLE transactions (
	id VARCHAR(64) NOT NULL PRIMARY KEY,
	account_id VARCHAR(64) NOT NULL REFERENCES accounts (id),
	creator VARCHAR(32) NOT NULL,
	type INTEGER NOT NULL,
	ref BLOB,
	status INTEGER NOT NULL,
	amount INTEGER NOT NULL,
	approval_date TIMESTAMP,
	approved_by VARCHAR(32),
	payment_date TIMESTAMP,
	creation_date TIMESTAMP,
	rejected_date TIMESTAMP,
	notes TEXT
)
