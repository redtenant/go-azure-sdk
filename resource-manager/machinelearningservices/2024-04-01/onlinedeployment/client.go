package onlinedeployment

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnlineDeploymentClient struct {
	Client *resourcemanager.Client
}

func NewOnlineDeploymentClientWithBaseURI(sdkApi sdkEnv.Api) (*OnlineDeploymentClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "onlinedeployment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnlineDeploymentClient: %+v", err)
	}

	return &OnlineDeploymentClient{
		Client: client,
	}, nil
}
