package securitysettings

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecuritySettingsClient struct {
	Client *resourcemanager.Client
}

func NewSecuritySettingsClientWithBaseURI(sdkApi sdkEnv.Api) (*SecuritySettingsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "securitysettings", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SecuritySettingsClient: %+v", err)
	}

	return &SecuritySettingsClient{
		Client: client,
	}, nil
}
