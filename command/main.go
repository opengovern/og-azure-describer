/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/kaytu-io/kaytu-azure-describer/azure"
	"github.com/kaytu-io/kaytu-util/pkg/describe/enums"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
)

var (
	resourceType, tenantID, clientID, clientSecret, subscriptionID string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kaytu-azure-describer",
	Short: "kaytu azure describer manual",
	RunE: func(cmd *cobra.Command, args []string) error {
		resourceType = "Microsoft.Resources/groups"
		subscriptionID = "710e21af-6987-4f5d-80a0-d2ef06f8645b"
		logger, _ := zap.NewProduction()
		output, err := azure.GetResources(
			context.Background(),
			logger,
			resourceType,
			enums.DescribeTriggerTypeManual,
			[]string{subscriptionID},
			azure.AuthConfig{
				TenantID:            "4725ad3d-5ab0-4f42-8a4a-fdee5ef586c5",     // tenantID,
				ClientID:            "08618331-3f87-4d97-bbe0-e3c4f06ae3cb",     // clientID,
				ClientSecret:        "c3~8Q~LDreBuHUwEAQGho2zLF1mcskeU3L4C-agy", // clientSecret,
				CertificatePath:     "",
				CertificatePassword: "",
				Username:            "",
				Password:            "",
			},
			string(azure.AuthEnv),
			"",
			nil,
		)
		if err != nil {
			return fmt.Errorf("Azure: %w", err)
		}
		js, err := json.Marshal(output)
		if err != nil {
			return err
		}
		fmt.Println(string(js))
		return nil
	},
}

func init() {
	rootCmd.Flags().StringVarP(&resourceType, "resourceType", "t", "", "Resource type")
	rootCmd.Flags().StringVarP(&tenantID, "tenantID", "", "", "TenantID")
	rootCmd.Flags().StringVarP(&clientID, "clientID", "", "", "ClientID")
	rootCmd.Flags().StringVarP(&clientSecret, "clientSecret", "", "", "ClientSecret")
	rootCmd.Flags().StringVarP(&subscriptionID, "subscriptionID", "", "", "SubscriptionID")
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
