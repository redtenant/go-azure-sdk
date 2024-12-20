package manageddeviceuser

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceUserClient struct {
	Client *msgraph.Client
}

func NewManagedDeviceUserClientWithBaseURI(sdkApi sdkEnv.Api) (*ManagedDeviceUserClient, error) {
	client, err := msgraph.NewClient(sdkApi, "manageddeviceuser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedDeviceUserClient: %+v", err)
	}

	return &ManagedDeviceUserClient{
		Client: client,
	}, nil
}
