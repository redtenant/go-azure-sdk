package vmmservers

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VMmServersClient struct {
	Client *resourcemanager.Client
}

func NewVMmServersClientWithBaseURI(sdkApi sdkEnv.Api) (*VMmServersClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "vmmservers", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VMmServersClient: %+v", err)
	}

	return &VMmServersClient{
		Client: client,
	}, nil
}
