package trustedidproviders

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TrustedIdProvidersClient struct {
	Client *resourcemanager.Client
}

func NewTrustedIdProvidersClientWithBaseURI(sdkApi sdkEnv.Api) (*TrustedIdProvidersClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "trustedidproviders", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TrustedIdProvidersClient: %+v", err)
	}

	return &TrustedIdProvidersClient{
		Client: client,
	}, nil
}
