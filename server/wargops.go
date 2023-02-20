package server

var Config config

type config struct {
	FrontendURI string `json:"frontend_uri"`
	// Discord
	DiscordGuildID      string `json:"discord_guild_id"`
	DiscordBotToken     string `json:"discord_bot_token"`
	DiscordClientID     string `json:"discord_client_id"`
	DiscordClientSecret string `json:"discord_client_secret"`
	DiscordRedirectURL  string `json:"discord_redirect_url"`
	LeadRoleID          string `json:"lead_role_id"`
	Secret              string `json:"secret"`
}

var (
	TypeDeposit       = TransactionType{0}
	TypeReimbursement = TransactionType{1}
	TypeProcurement   = TransactionType{2}

	StatusCreated  = TransactionStatus{0}
	StatusPending  = TransactionStatus{1}
	StatusPaid     = TransactionStatus{2}
	StatusRejected = TransactionStatus{3}
)
