package claimspolicy

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClaimsPolicyClient struct {
	Client *msgraph.Client
}

func NewClaimsPolicyClientWithBaseURI(sdkApi sdkEnv.Api) (*ClaimsPolicyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "claimspolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ClaimsPolicyClient: %+v", err)
	}

	return &ClaimsPolicyClient{
		Client: client,
	}, nil
}
