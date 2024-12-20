package hardwareconfigurationdevicerunstate

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HardwareConfigurationDeviceRunStateClient struct {
	Client *msgraph.Client
}

func NewHardwareConfigurationDeviceRunStateClientWithBaseURI(sdkApi sdkEnv.Api) (*HardwareConfigurationDeviceRunStateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "hardwareconfigurationdevicerunstate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating HardwareConfigurationDeviceRunStateClient: %+v", err)
	}

	return &HardwareConfigurationDeviceRunStateClient{
		Client: client,
	}, nil
}
