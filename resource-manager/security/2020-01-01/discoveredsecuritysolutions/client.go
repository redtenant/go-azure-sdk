package discoveredsecuritysolutions

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiscoveredSecuritySolutionsClient struct {
	Client *resourcemanager.Client
}

func NewDiscoveredSecuritySolutionsClientWithBaseURI(sdkApi sdkEnv.Api) (*DiscoveredSecuritySolutionsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "discoveredsecuritysolutions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DiscoveredSecuritySolutionsClient: %+v", err)
	}

	return &DiscoveredSecuritySolutionsClient{
		Client: client,
	}, nil
}
