package virtualendpointprovisioningpolicy

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEndpointProvisioningPolicyClient struct {
	Client *msgraph.Client
}

func NewVirtualEndpointProvisioningPolicyClientWithBaseURI(sdkApi sdkEnv.Api) (*VirtualEndpointProvisioningPolicyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "virtualendpointprovisioningpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VirtualEndpointProvisioningPolicyClient: %+v", err)
	}

	return &VirtualEndpointProvisioningPolicyClient{
		Client: client,
	}, nil
}
