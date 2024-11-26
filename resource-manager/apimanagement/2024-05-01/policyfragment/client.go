package policyfragment

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyFragmentClient struct {
	Client *resourcemanager.Client
}

func NewPolicyFragmentClientWithBaseURI(sdkApi sdkEnv.Api) (*PolicyFragmentClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "policyfragment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PolicyFragmentClient: %+v", err)
	}

	return &PolicyFragmentClient{
		Client: client,
	}, nil
}
