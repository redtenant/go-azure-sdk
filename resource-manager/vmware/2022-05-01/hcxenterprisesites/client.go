package hcxenterprisesites

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HcxEnterpriseSitesClient struct {
	Client *resourcemanager.Client
}

func NewHcxEnterpriseSitesClientWithBaseURI(sdkApi sdkEnv.Api) (*HcxEnterpriseSitesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "hcxenterprisesites", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating HcxEnterpriseSitesClient: %+v", err)
	}

	return &HcxEnterpriseSitesClient{
		Client: client,
	}, nil
}
