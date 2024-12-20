package registries

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegistriesClient struct {
	Client *resourcemanager.Client
}

func NewRegistriesClientWithBaseURI(sdkApi sdkEnv.Api) (*RegistriesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "registries", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RegistriesClient: %+v", err)
	}

	return &RegistriesClient{
		Client: client,
	}, nil
}
