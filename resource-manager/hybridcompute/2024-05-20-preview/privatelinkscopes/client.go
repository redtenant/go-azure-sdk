package privatelinkscopes

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateLinkScopesClient struct {
	Client *resourcemanager.Client
}

func NewPrivateLinkScopesClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivateLinkScopesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "privatelinkscopes", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivateLinkScopesClient: %+v", err)
	}

	return &PrivateLinkScopesClient{
		Client: client,
	}, nil
}
