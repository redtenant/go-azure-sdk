package serveroperations

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerOperationsClient struct {
	Client *resourcemanager.Client
}

func NewServerOperationsClientWithBaseURI(sdkApi sdkEnv.Api) (*ServerOperationsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "serveroperations", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ServerOperationsClient: %+v", err)
	}

	return &ServerOperationsClient{
		Client: client,
	}, nil
}
