package clusterextensions

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClusterExtensionsClient struct {
	Client *resourcemanager.Client
}

func NewClusterExtensionsClientWithBaseURI(sdkApi sdkEnv.Api) (*ClusterExtensionsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "clusterextensions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ClusterExtensionsClient: %+v", err)
	}

	return &ClusterExtensionsClient{
		Client: client,
	}, nil
}
