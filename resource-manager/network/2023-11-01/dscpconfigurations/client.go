package dscpconfigurations

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DscpConfigurationsClient struct {
	Client *resourcemanager.Client
}

func NewDscpConfigurationsClientWithBaseURI(sdkApi sdkEnv.Api) (*DscpConfigurationsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "dscpconfigurations", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DscpConfigurationsClient: %+v", err)
	}

	return &DscpConfigurationsClient{
		Client: client,
	}, nil
}
