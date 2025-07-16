package helpers

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

// Verbose controls verbose logging
var Verbose bool

func LogVerbose(format string, a ...interface{}) {
	if Verbose {
		fmt.Printf(format, a...)
	}
}

func LoadAwsConfig() (aws.Config, string) {

	profile, ok := os.LookupEnv("aws_profile")

	if !ok {
		profile = "default"
	}

	if profile == "" {
		log.Fatal("Need a profile to continue")
	}

	// Load AWS configuration from default config sources
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile(profile))

	// Handle any config loading errors
	if err != nil {
		fmt.Printf("❌ Error: failed to load AWS config: %v\n", err)
		panic(err)
	}

	LogVerbose("🔑 Using AWS profile: %s\n", profile)
	LogVerbose("🌎 Using AWS region: %s\n", cfg.Region)

	// Create a new STS client using the loaded config
	client := sts.NewFromConfig(cfg)

	// Get the caller identity to validate credentials and get account info
	result, err := client.GetCallerIdentity(context.TODO(), &sts.GetCallerIdentityInput{})
	if err != nil {
		log.Fatalf("❌ Error: failed to get caller identity: %v\n", err)
	}

	// Print account and identity information
	LogVerbose("🏦 Using AWS account ID: %s\n", *result.Account)
	LogVerbose("🙍 Using caller identity: %s\n", *result.Arn)

	return cfg, *result.Account
}
