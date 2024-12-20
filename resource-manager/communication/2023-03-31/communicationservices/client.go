package communicationservices

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommunicationServicesClient struct {
	Client *resourcemanager.Client
}

func NewCommunicationServicesClientWithBaseURI(sdkApi sdkEnv.Api) (*CommunicationServicesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "communicationservices", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CommunicationServicesClient: %+v", err)
	}

	return &CommunicationServicesClient{
		Client: client,
	}, nil
}
