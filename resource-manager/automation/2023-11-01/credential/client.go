package credential

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CredentialClient struct {
	Client *resourcemanager.Client
}

func NewCredentialClientWithBaseURI(sdkApi sdkEnv.Api) (*CredentialClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "credential", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CredentialClient: %+v", err)
	}

	return &CredentialClient{
		Client: client,
	}, nil
}
