package servicerunners

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceRunnersClient struct {
	Client *resourcemanager.Client
}

func NewServiceRunnersClientWithBaseURI(sdkApi sdkEnv.Api) (*ServiceRunnersClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "servicerunners", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ServiceRunnersClient: %+v", err)
	}

	return &ServiceRunnersClient{
		Client: client,
	}, nil
}
