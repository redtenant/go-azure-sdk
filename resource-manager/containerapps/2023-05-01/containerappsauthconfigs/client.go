package containerappsauthconfigs

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContainerAppsAuthConfigsClient struct {
	Client *resourcemanager.Client
}

func NewContainerAppsAuthConfigsClientWithBaseURI(sdkApi sdkEnv.Api) (*ContainerAppsAuthConfigsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "containerappsauthconfigs", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ContainerAppsAuthConfigsClient: %+v", err)
	}

	return &ContainerAppsAuthConfigsClient{
		Client: client,
	}, nil
}
