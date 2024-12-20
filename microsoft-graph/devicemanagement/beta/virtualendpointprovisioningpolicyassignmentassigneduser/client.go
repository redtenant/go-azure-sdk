package virtualendpointprovisioningpolicyassignmentassigneduser

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEndpointProvisioningPolicyAssignmentAssignedUserClient struct {
	Client *msgraph.Client
}

func NewVirtualEndpointProvisioningPolicyAssignmentAssignedUserClientWithBaseURI(sdkApi sdkEnv.Api) (*VirtualEndpointProvisioningPolicyAssignmentAssignedUserClient, error) {
	client, err := msgraph.NewClient(sdkApi, "virtualendpointprovisioningpolicyassignmentassigneduser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VirtualEndpointProvisioningPolicyAssignmentAssignedUserClient: %+v", err)
	}

	return &VirtualEndpointProvisioningPolicyAssignmentAssignedUserClient{
		Client: client,
	}, nil
}
