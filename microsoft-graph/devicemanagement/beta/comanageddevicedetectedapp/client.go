package comanageddevicedetectedapp

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComanagedDeviceDetectedAppClient struct {
	Client *msgraph.Client
}

func NewComanagedDeviceDetectedAppClientWithBaseURI(sdkApi sdkEnv.Api) (*ComanagedDeviceDetectedAppClient, error) {
	client, err := msgraph.NewClient(sdkApi, "comanageddevicedetectedapp", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ComanagedDeviceDetectedAppClient: %+v", err)
	}

	return &ComanagedDeviceDetectedAppClient{
		Client: client,
	}, nil
}
