package describer

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"

	"github.com/kaytu-io/kaytu-util/pkg/describe"
	"github.com/kaytu-io/kaytu-util/pkg/vault"
	"github.com/kaytu-io/kaytu-util/proto/src/golang"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
	"google.golang.org/grpc/metadata"
)

const (
	DescribeResourceJobFailed    string = "FAILED"
	DescribeResourceJobSucceeded string = "SUCCEEDED"
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
		"https://app.kaytu.io/workspaceAccess": map[string]string{
			workspaceId: "admin",
		},
		"https://app.kaytu.io/email": "lambda-worker@kaytu.io",
	}).SignedString(pk)
	if err != nil {
		return "", fmt.Errorf("JWT token generation failed %v", err)
	}
	return token, nil
}

type TriggeredBy string

const (
	TriggeredByAWSLambda     TriggeredBy = "aws-lambda"
	TriggeredByAzureFunction TriggeredBy = "azure-function"
)

// DescribeHandler
// TriggeredBy is not used for now but might be relevant in the future
func DescribeHandler(ctx context.Context, logger *zap.Logger, _ TriggeredBy, input describe.DescribeWorkerInput) error {
	var err error

	if input.WorkspaceName == "" {
		return fmt.Errorf("workspace name is required")
	}

	token, err := getJWTAuthToken(input.WorkspaceId)
	if err != nil {
		return fmt.Errorf("failed to get JWT token: %w", err)
	}

	var client golang.DescribeServiceClient
	grpcCtx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{
		"workspace-name": input.WorkspaceName,
	}))
	for retry := 0; retry < 5; retry++ {
		conn, err := grpc.Dial(
			input.DescribeEndpoint,
			grpc.WithTransportCredentials(credentials.NewTLS(nil)),
			grpc.WithPerRPCCredentials(oauth.TokenSource{
				TokenSource: oauth2.StaticTokenSource(&oauth2.Token{
					AccessToken: token,
				}),
			}),
		)
		if err != nil {
			logger.Error("[result delivery] connection failure:", zap.Error(err))
			if retry == 4 {
				return err
			}
			time.Sleep(1 * time.Second)
			continue
		}
		client = golang.NewDescribeServiceClient(conn)
		break
	}

	for retry := 0; retry < 5; retry++ {
		_, err := client.SetInProgress(grpcCtx, &golang.SetInProgressRequest{
			JobId: uint32(input.DescribeJob.JobID),
		})
		if err != nil {
			logger.Error("[result delivery] set in progress failure:", zap.Error(err))
			if retry == 4 {
				return err
			}
			time.Sleep(1 * time.Second)
			continue
		}
		break
	}

	var vaultSc vault.VaultSourceConfig
	switch input.VaultConfig.Provider {
	case vault.AwsKMS:
		vaultSc, err = vault.NewKMSVaultSourceConfig(ctx, input.VaultConfig.Aws, input.VaultConfig.KeyId)
		if err != nil {
			return fmt.Errorf("failed to initialize KMS vault: %w", err)
		}
	case vault.AzureKeyVault:
		vaultSc, err = vault.NewAzureVaultClient(ctx, logger, input.VaultConfig.Azure, input.VaultConfig.KeyId)
		if err != nil {
			return fmt.Errorf("failed to initialize Azure vault: %w", err)
		}
	}

	resourceIds, err := Do(
		ctx,
		vaultSc,
		logger,
		input.DescribeJob,
		input.VaultConfig.KeyId,
		input.DescribeEndpoint,
		token,
		input.IngestionPipelineEndpoint,
		input.UseOpenSearch,
		input.WorkspaceName,
		input.WorkspaceId,
	)

	errMsg := ""
	errCode := ""
	status := DescribeResourceJobSucceeded
	if err != nil {
		//errMsg = err.Error()
		//var detailedErr autorest.DetailedError
		//if errors.As(err, &detailedErr) {
		//	errCode = fmt.Sprintf("%v", detailedErr.StatusCode)
		//}
		//
		//var validationErr validation.Error
		//if errors.As(err, &validationErr) {
		//	errCode = "ValidationError"
		//}
		//
		//if errCode == "" {
		//	if strings.Contains(err.Error(), "InvalidAuthenticationToken") {
		//		errCode = "InvalidAuthenticationToken"
		//	}
		//}

		errorString := err.Error()
		jsonStart := 0
		jsonEnd := len(errorString) - 1
		for i := 0; i < len(errorString)-1; i++ {
			if errorString[i] == '{' {
				jsonStart = i
				break
			}
		}
		for i := len(errorString) - 1; i > 0; i-- {
			if errorString[i] == '}' {
				jsonEnd = i + 1
				break
			}
		}

		jsonString := errorString[jsonStart:jsonEnd]
		var jsonData map[string]interface{}
		err := json.Unmarshal([]byte(jsonString), &jsonData)
		if err != nil || jsonData == nil {
			errMsg = errorString
			errCode = "UnknownFailure"
		} else {
			if errorData, ok := jsonData["error"]; ok {
				if v, ok := errorData.(map[string]interface{}); ok {
					if code, ok := v["code"].(string); ok {
						errCode = code
					} else {
						errCode = "unknown"
					}
					if msg, ok := v["message"].(string); ok {
						errMsg = msg
					} else {
						errMsg = fmt.Sprintf("ErrMsg= %v", v)
					}
				} else {
					errCode = "unknown"
					errMsg = fmt.Sprintf("ErrorData= %v", errorData)
				}
			} else {
				errMsg = errorString
				errCode = "UnknownFailure"
			}
		}
		status = DescribeResourceJobFailed
	}

	for retry := 0; retry < 5; retry++ {
		_, err = client.DeliverResult(grpcCtx, &golang.DeliverResultRequest{
			JobId:     uint32(input.DescribeJob.JobID),
			Status:    status,
			Error:     errMsg,
			ErrorCode: errCode,
			DescribeJob: &golang.DescribeJob{
				JobId:        uint32(input.DescribeJob.JobID),
				ResourceType: input.DescribeJob.ResourceType,
				SourceId:     input.DescribeJob.SourceID,
				AccountId:    input.DescribeJob.AccountID,
				DescribedAt:  input.DescribeJob.DescribedAt,
				SourceType:   string(input.DescribeJob.SourceType),
				ConfigReg:    input.DescribeJob.CipherText,
				TriggerType:  string(input.DescribeJob.TriggerType),
				RetryCounter: uint32(input.DescribeJob.RetryCounter),
			},
			DescribedResourceIds: resourceIds,
		})
		if err != nil {
			logger.Error("[result delivery] rpc failed:", zap.Error(err))
			time.Sleep(1 * time.Second)
			continue
		}
		break
	}

	return err
}
