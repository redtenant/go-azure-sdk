package partnerconfigurations

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PartnerConfigurationsClient struct {
	Client *resourcemanager.Client
}

func NewPartnerConfigurationsClientWithBaseURI(sdkApi sdkEnv.Api) (*PartnerConfigurationsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "partnerconfigurations", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PartnerConfigurationsClient: %+v", err)
	}

	return &PartnerConfigurationsClient{
		Client: client,
	}, nil
}
