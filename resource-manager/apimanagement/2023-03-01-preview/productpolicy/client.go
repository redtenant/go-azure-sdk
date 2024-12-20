package productpolicy

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProductPolicyClient struct {
	Client *resourcemanager.Client
}

func NewProductPolicyClientWithBaseURI(sdkApi sdkEnv.Api) (*ProductPolicyClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "productpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ProductPolicyClient: %+v", err)
	}

	return &ProductPolicyClient{
		Client: client,
	}, nil
}
