package devicecommandresponsepayload

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCommandResponsepayloadClient struct {
	Client *msgraph.Client
}

func NewDeviceCommandResponsepayloadClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceCommandResponsepayloadClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicecommandresponsepayload", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceCommandResponsepayloadClient: %+v", err)
	}

	return &DeviceCommandResponsepayloadClient{
		Client: client,
	}, nil
}
