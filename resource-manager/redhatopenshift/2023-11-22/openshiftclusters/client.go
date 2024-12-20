package openshiftclusters

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OpenShiftClustersClient struct {
	Client *resourcemanager.Client
}

func NewOpenShiftClustersClientWithBaseURI(sdkApi sdkEnv.Api) (*OpenShiftClustersClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "openshiftclusters", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OpenShiftClustersClient: %+v", err)
	}

	return &OpenShiftClustersClient{
		Client: client,
	}, nil
}
