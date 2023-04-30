package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/golang-jwt/jwt/v5"
	kaytu_azure_describer "github.com/kaytu-io/kaytu-azure-describer"
	"github.com/kaytu-io/kaytu-azure-describer/pkg/describe"
	"github.com/kaytu-io/kaytu-azure-describer/pkg/vault"
	"go.uber.org/zap"
)

func getJWTAuthToken(workspaceId string) (string, error) {
	privateKey, ok := os.LookupEnv("JWT_PRIVATE_KEY")
	if !ok {
		return "", fmt.Errorf("JWT_PRIVATE_KEY not set")
	}

	privateKeyBytes, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return "", fmt.Errorf("JWT_PRIVATE_KEY not base64 encoded")
	}

	pk, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		return "", fmt.Errorf("JWT_PRIVATE_KEY not valid")
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"https://app.keibi.io/workspaceAccess": map[string]string{
			workspaceId: "admin",
		},
		"https://app.keibi.io/email": "lambda-worker@kaytu.io",
	}).SignedString(pk)
	if err != nil {
		return "", fmt.Errorf("JWT token generation failed %v", err)
	}
	return token, nil
}

func DescribeHandler(ctx context.Context, input describe.LambdaDescribeWorkerInput) error {
	logger, err := zap.NewProduction()
	if err != nil {
		return err
	}

	logger.Info(fmt.Sprintf("%v", input))

	kmsVault, err := vault.NewKMSVaultSourceConfig(ctx)
	if err != nil {
		return fmt.Errorf("failed to initialize KMS vault: %w", err)
	}

	token, err := getJWTAuthToken(input.WorkspaceId)
	if err != nil {
		return fmt.Errorf("failed to get JWT token: %w", err)
	}

	return kaytu_azure_describer.Do(
		ctx,
		kmsVault,
		input.DescribeJob,
		input.KeyARN,
		logger,
		&input.DescribeEndpoint,
		&token,
	)
}

func main() {
	lambda.Start(DescribeHandler)
}
