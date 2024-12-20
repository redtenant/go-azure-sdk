package registereddevice

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegisteredDeviceClient struct {
	Client *msgraph.Client
}

func NewRegisteredDeviceClientWithBaseURI(sdkApi sdkEnv.Api) (*RegisteredDeviceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "registereddevice", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RegisteredDeviceClient: %+v", err)
	}

	return &RegisteredDeviceClient{
		Client: client,
	}, nil
}
