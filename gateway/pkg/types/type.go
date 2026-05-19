package types




type CreateWalletRequest struct {
	WorkspaceID string `json:"workspace_id"`
}

type TransactionRequest struct {
	WorkspaceID string `json:"workspace_id"`
	ToAddress   string `json:"to_address"`
	Amount      string `json:"amount"`
	Nonce       uint64 `json:"nonce"`
}
