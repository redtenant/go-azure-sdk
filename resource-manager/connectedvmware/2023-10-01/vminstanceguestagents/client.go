package vminstanceguestagents

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VMInstanceGuestAgentsClient struct {
	Client *resourcemanager.Client
}

func NewVMInstanceGuestAgentsClientWithBaseURI(sdkApi sdkEnv.Api) (*VMInstanceGuestAgentsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "vminstanceguestagents", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VMInstanceGuestAgentsClient: %+v", err)
	}

	return &VMInstanceGuestAgentsClient{
		Client: client,
	}, nil
}
