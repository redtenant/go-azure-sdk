package workloadclassifiers

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkloadClassifiersClient struct {
	Client *resourcemanager.Client
}

func NewWorkloadClassifiersClientWithBaseURI(sdkApi sdkEnv.Api) (*WorkloadClassifiersClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "workloadclassifiers", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WorkloadClassifiersClient: %+v", err)
	}

	return &WorkloadClassifiersClient{
		Client: client,
	}, nil
}
