package guestconfigurationconnectedvmwarevsphereassignmentsreports

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GuestConfigurationConnectedVMwarevSphereAssignmentsReportsClient struct {
	Client *resourcemanager.Client
}

func NewGuestConfigurationConnectedVMwarevSphereAssignmentsReportsClientWithBaseURI(sdkApi sdkEnv.Api) (*GuestConfigurationConnectedVMwarevSphereAssignmentsReportsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "guestconfigurationconnectedvmwarevsphereassignmentsreports", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GuestConfigurationConnectedVMwarevSphereAssignmentsReportsClient: %+v", err)
	}

	return &GuestConfigurationConnectedVMwarevSphereAssignmentsReportsClient{
		Client: client,
	}, nil
}
