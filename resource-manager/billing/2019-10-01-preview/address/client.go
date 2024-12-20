package address

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddressClient struct {
	Client *resourcemanager.Client
}

func NewAddressClientWithBaseURI(sdkApi sdkEnv.Api) (*AddressClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "address", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AddressClient: %+v", err)
	}

	return &AddressClient{
		Client: client,
	}, nil
}
