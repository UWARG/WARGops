package main

var (
	TypeDeposit       = TransactionType{0}
	TypeReimbursement = TransactionType{1}
	TypeProcurement   = TransactionType{2}

	StatusCreated  = TransactionStatus{0}
	StatusPending  = TransactionStatus{1}
	StatusPaid     = TransactionStatus{2}
	StatusRejected = TransactionStatus{3}
)
