package azure

type AzureADCredentials struct {
	SubscriptionID string `json:"subscription"`
	TenantID       string `json:"tenant_id"`
	ClientID       string `json:"client_id"`
	ClientSecret   string `json:"client_secret"`
}

// func (aad *AzureADCredentials) GetCredentials() {

// }
