package virtualendpointsnapshot

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEndpointSnapshotClient struct {
	Client *msgraph.Client
}

func NewVirtualEndpointSnapshotClientWithBaseURI(sdkApi sdkEnv.Api) (*VirtualEndpointSnapshotClient, error) {
	client, err := msgraph.NewClient(sdkApi, "virtualendpointsnapshot", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VirtualEndpointSnapshotClient: %+v", err)
	}

	return &VirtualEndpointSnapshotClient{
		Client: client,
	}, nil
}
