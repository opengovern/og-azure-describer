package vault

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
)

type KMSVaultSourceConfig struct {
	kmsClient *kms.Client
}

func NewKMSVaultSourceConfig(ctx context.Context) (*KMSVaultSourceConfig, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load SDK configuration: %v", err)
	}

	// Create KMS client with loaded configuration
	svc := kms.NewFromConfig(cfg)

	return &KMSVaultSourceConfig{
		kmsClient: svc,
	}, nil
}

func (v *KMSVaultSourceConfig) Encrypt(cred map[string]any, keyARN string) ([]byte, error) {
	bytes, err := json.Marshal(cred)
	if err != nil {
		return nil, err
	}

	result, err := v.kmsClient.Encrypt(context.TODO(), &kms.EncryptInput{
		KeyId:               &keyARN,
		Plaintext:           bytes,
		EncryptionAlgorithm: types.EncryptionAlgorithmSpecSymmetricDefault,
		EncryptionContext:   nil, //TODO-Saleh use workspaceID
		GrantTokens:         nil,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to encrypt ciphertext: %v", err)
	}
	return result.CiphertextBlob, nil
}

func (v *KMSVaultSourceConfig) Decrypt(cypherText string, keyARN string) (map[string]any, error) {
	result, err := v.kmsClient.Decrypt(context.TODO(), &kms.DecryptInput{
		CiphertextBlob:      []byte(cypherText),
		EncryptionAlgorithm: types.EncryptionAlgorithmSpecSymmetricDefault,
		KeyId:               &keyARN,
		EncryptionContext:   nil, //TODO-Saleh use workspaceID
	})
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt ciphertext: %v", err)
	}

	conf := make(map[string]any)
	err = json.Unmarshal(result.Plaintext, &conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}
