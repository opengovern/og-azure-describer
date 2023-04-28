package azure

import (
	"encoding/json"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

type AuthConfig struct {
	TenantID            string
	ObjectID            string
	SecretID            string
	AuxiliaryTenantIDs  string
	ClientID            string
	ClientSecret        string
	CertificatePath     string
	CertificatePassword string
	Username            string
	Password            string
	EnvironmentName     string
	Resource            string
}

func NewAuthorizerFromConfig(cfg AuthConfig) (autorest.Authorizer, error) {
	settings, err := GetSettingsFromConfig(cfg)
	if err != nil {
		return nil, err
	}
	return settings.GetAuthorizer()
}

func GetSettingsFromConfig(cfg AuthConfig) (s auth.EnvironmentSettings, err error) {
	s = auth.EnvironmentSettings{
		Values: map[string]string{},
	}

	if cfg.TenantID != "" {
		s.Values[auth.TenantID] = cfg.TenantID
	}
	if cfg.AuxiliaryTenantIDs != "" {
		s.Values[auth.AuxiliaryTenantIDs] = cfg.AuxiliaryTenantIDs
	}
	if cfg.ClientID != "" {
		s.Values[auth.ClientID] = cfg.ClientID
	}
	if cfg.ClientSecret != "" {
		s.Values[auth.ClientSecret] = cfg.ClientSecret
	}
	if cfg.CertificatePath != "" {
		s.Values[auth.CertificatePath] = cfg.CertificatePath
	}
	if cfg.CertificatePassword != "" {
		s.Values[auth.CertificatePassword] = cfg.CertificatePassword
	}
	if cfg.Username != "" {
		s.Values[auth.Username] = cfg.Username
	}
	if cfg.Password != "" {
		s.Values[auth.Password] = cfg.Password
	}
	if cfg.EnvironmentName != "" {
		s.Values[auth.EnvironmentName] = cfg.EnvironmentName
	}
	if cfg.Resource != "" {
		s.Values[auth.Resource] = cfg.Resource
	}

	if v := s.Values[auth.EnvironmentName]; v == "" {
		s.Environment = azure.PublicCloud
	} else {
		s.Environment, err = azure.EnvironmentFromName(v)
	}
	if s.Values[auth.Resource] == "" {
		s.Values[auth.Resource] = s.Environment.ResourceManagerEndpoint
	}
	return
}

type SubscriptionConfig struct {
	SubscriptionID  string `json:"subscriptionId"`
	TenantID        string `json:"tenantId"`
	ObjectID        string `json:"objectId"`
	SecretID        string `json:"secretId"`
	ClientID        string `json:"clientId"`
	ClientSecret    string `json:"clientSecret"`
	CertificatePath string `json:"certificatePath"`
	CertificatePass string `json:"certificatePass"`
	Username        string `json:"username"`
	Password        string `json:"password"`
}

func SubscriptionConfigFromMap(m map[string]any) (SubscriptionConfig, error) {
	mj, err := json.Marshal(m)
	if err != nil {
		return SubscriptionConfig{}, err
	}

	var c SubscriptionConfig
	err = json.Unmarshal(mj, &c)
	if err != nil {
		return SubscriptionConfig{}, err
	}

	return c, nil
}
