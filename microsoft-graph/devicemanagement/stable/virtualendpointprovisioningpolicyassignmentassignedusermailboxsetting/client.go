package virtualendpointprovisioningpolicyassignmentassignedusermailboxsetting

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEndpointProvisioningPolicyAssignmentAssignedUserMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewVirtualEndpointProvisioningPolicyAssignmentAssignedUserMailboxSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*VirtualEndpointProvisioningPolicyAssignmentAssignedUserMailboxSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "virtualendpointprovisioningpolicyassignmentassignedusermailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VirtualEndpointProvisioningPolicyAssignmentAssignedUserMailboxSettingClient: %+v", err)
	}

	return &VirtualEndpointProvisioningPolicyAssignmentAssignedUserMailboxSettingClient{
		Client: client,
	}, nil
}
