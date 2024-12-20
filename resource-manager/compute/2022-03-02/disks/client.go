package disks

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DisksClient struct {
	Client *resourcemanager.Client
}

func NewDisksClientWithBaseURI(sdkApi sdkEnv.Api) (*DisksClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "disks", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DisksClient: %+v", err)
	}

	return &DisksClient{
		Client: client,
	}, nil
}
