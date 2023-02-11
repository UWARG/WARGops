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

export interface NewTransaction{
    id: string,
    name: string,
    account_id: string,
    amount: number,
    type: number,
    status: number,
    notes: string,
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
    name: string,
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
    accent_color: string,
    avatar: string,
    avatar_decoration: null,
    banner: null,
    banner_color: "#6b4f31",
    discriminator: "3963",
    display_name: null,
    flags: 0,
    id: string,
    locale: "en-US",
    mfa_enabled: false,
    premium_type: 0,
    public_flags: 0,
    username: string;
    role: "lead" | "bootcamper" | "dev";
}
