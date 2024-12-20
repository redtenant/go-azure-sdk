package securityrules

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityRulesClient struct {
	Client *resourcemanager.Client
}

func NewSecurityRulesClientWithBaseURI(sdkApi sdkEnv.Api) (*SecurityRulesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "securityrules", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SecurityRulesClient: %+v", err)
	}

	return &SecurityRulesClient{
		Client: client,
	}, nil
}
