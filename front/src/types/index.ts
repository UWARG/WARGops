export interface Account {
    name: 'WEEF',
    id: string,
    source: string,
    active: true,
    allocation_date: Date,
    expiry_ate: Date,
    creator: string,
    point_of_contact: string, // must be email
    waterloo_id: string;
}

export interface NewTransaction extends Transaction {
    id: string,
    account_id: string,
    amount: number,
    type: number,
    status: number,
}


// Implement: 
// AccountID    string            `json:"account_id"`
// 	Amount       int               `json:"amount"`
// 	ApprovalDate time.Time         `json:"approval_date"`
// 	CreationDate time.Time         `json:"creation_date"`
// 	ID           string            `json:"id"`
// 	Notes        string            `json:"notes"`
// 	PaymentDate  time.Time         `json:"payment_date"`
// 	RejectedDate time.Time         `json:"rejected_date"`
// 	Status       TransactionStatus `json:"status"`
// 	Type         TransactionType   `json:"type"`

export interface Transaction {
    id: string,
    amount: number,
    approval_date: Date,
    creation_date: Date,
    notes: string,
    payment_date: Date,
    rejected_date: Date,
    status: number,
    type: number,
}

export interface Profile {
    code: string,
    access_token: string,
    refresh_token: string,
    expires_in: number,
    token_type: string,
    scope: string,
    id_token: string,
    user_id: string,
    user_name: string,
    user_email: string,
    user_avatar: string,
    user_discriminator: string,
    user_public_flags: number,
    user_flags: number,
    user_locale: string,
}