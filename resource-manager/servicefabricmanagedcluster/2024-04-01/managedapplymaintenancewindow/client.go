package managedapplymaintenancewindow

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedApplyMaintenanceWindowClient struct {
	Client *resourcemanager.Client
}

func NewManagedApplyMaintenanceWindowClientWithBaseURI(sdkApi sdkEnv.Api) (*ManagedApplyMaintenanceWindowClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "managedapplymaintenancewindow", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedApplyMaintenanceWindowClient: %+v", err)
	}

	return &ManagedApplyMaintenanceWindowClient{
		Client: client,
	}, nil
}
