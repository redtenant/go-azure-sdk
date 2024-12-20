package tenantsettings

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TenantSettingsClient struct {
	Client *resourcemanager.Client
}

func NewTenantSettingsClientWithBaseURI(sdkApi sdkEnv.Api) (*TenantSettingsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "tenantsettings", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TenantSettingsClient: %+v", err)
	}

	return &TenantSettingsClient{
		Client: client,
	}, nil
}
