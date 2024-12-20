package remotedesktopsecurityconfiguration

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemoteDesktopSecurityConfigurationClient struct {
	Client *msgraph.Client
}

func NewRemoteDesktopSecurityConfigurationClientWithBaseURI(sdkApi sdkEnv.Api) (*RemoteDesktopSecurityConfigurationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "remotedesktopsecurityconfiguration", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RemoteDesktopSecurityConfigurationClient: %+v", err)
	}

	return &RemoteDesktopSecurityConfigurationClient{
		Client: client,
	}, nil
}
