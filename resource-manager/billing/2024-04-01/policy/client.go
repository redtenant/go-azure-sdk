package policy

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyClient struct {
	Client *resourcemanager.Client
}

func NewPolicyClientWithBaseURI(sdkApi sdkEnv.Api) (*PolicyClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "policy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PolicyClient: %+v", err)
	}

	return &PolicyClient{
		Client: client,
	}, nil
}
