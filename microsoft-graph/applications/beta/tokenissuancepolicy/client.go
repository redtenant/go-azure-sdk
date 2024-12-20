package tokenissuancepolicy

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TokenIssuancePolicyClient struct {
	Client *msgraph.Client
}

func NewTokenIssuancePolicyClientWithBaseURI(sdkApi sdkEnv.Api) (*TokenIssuancePolicyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "tokenissuancepolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TokenIssuancePolicyClient: %+v", err)
	}

	return &TokenIssuancePolicyClient{
		Client: client,
	}, nil
}
