package comanageddeviceuser

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComanagedDeviceUserClient struct {
	Client *msgraph.Client
}

func NewComanagedDeviceUserClientWithBaseURI(sdkApi sdkEnv.Api) (*ComanagedDeviceUserClient, error) {
	client, err := msgraph.NewClient(sdkApi, "comanageddeviceuser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ComanagedDeviceUserClient: %+v", err)
	}

	return &ComanagedDeviceUserClient{
		Client: client,
	}, nil
}
