package helpers

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

func LoadAwsConfig() (aws.Config, string) {

	// Get AWS profile name from user selection
	profile := "default"

	// Validate that a profile was selected
	if profile == "" {
		log.Fatal("Need a profile to continue")
	}

	// Load AWS configuration from default config sources
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile(profile))

	// Handle any config loading errors
	if err != nil {
		panic(err)
	}

	fmt.Printf("🔑 Using AWS profile: %s\n", profile)
	fmt.Printf("🌎 Using AWS region: %s\n", cfg.Region)

	// Create a new STS client using the loaded config
	client := sts.NewFromConfig(cfg)

	// Get the caller identity to validate credentials and get account info
	result, err := client.GetCallerIdentity(context.TODO(), &sts.GetCallerIdentityInput{})
	if err != nil {
		panic("failed to get caller identity, " + err.Error())
	}

	// Print account and identity information
	fmt.Printf("🏦 Using AWS account ID: %s\n", *result.Account)
	fmt.Printf("🙍 Using caller identity: %s\n", *result.Arn)

	return cfg, *result.Account
}
