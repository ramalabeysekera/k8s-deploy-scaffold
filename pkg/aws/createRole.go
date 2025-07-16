package aws

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/ramalabeysekera/k8s-deploy-scaffold/helpers"
)

func CreateIamRole(config aws.Config, serviceAccountName string, accountId string, namespaceName string, oidcProvider string) (string, error) {

	client := iam.NewFromConfig(config)

	policy := map[string]interface{}{
		"Version": "2012-10-17",
		"Statement": []map[string]interface{}{
			{
				"Effect": "Allow",
				"Principal": map[string]string{
					"Federated": fmt.Sprintf("arn:aws:iam::%s:oidc-provider/%s", accountId, oidcProvider),
				},
				"Action": "sts:AssumeRoleWithWebIdentity",
				"Condition": map[string]map[string]string{
					"StringEquals": {
						fmt.Sprintf("%s:aud", oidcProvider): "sts.amazonaws.com",
						fmt.Sprintf("%s:sub", oidcProvider): fmt.Sprintf("system:serviceaccount:%s:%s", namespaceName, serviceAccountName),
					},
				},
			},
		},
	}

	policyBytes, err := json.Marshal(policy)
	if err != nil {
		return "", fmt.Errorf("failed to marshal policy: %w", err)
	}

	policyString := string(policyBytes)
	description := fmt.Sprintf("IAM role for %s service account", serviceAccountName)

	createRoleInput := iam.CreateRoleInput{
		AssumeRolePolicyDocument: &policyString,
		RoleName:                 &serviceAccountName,
		Description:              &description,
	}

	createRoleOutput, err := client.CreateRole(context.Background(), &createRoleInput)
	if err != nil {
		return "", err
	}

	helpers.LogVerbose("IAM policy: %s\n", policyString)
	helpers.LogVerbose("✅ Created role %s at %v\n", serviceAccountName, createRoleOutput.Role.CreateDate)
	return *createRoleOutput.Role.Arn, nil
}
