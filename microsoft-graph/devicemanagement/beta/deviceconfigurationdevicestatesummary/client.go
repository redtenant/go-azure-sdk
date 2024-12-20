package deviceconfigurationdevicestatesummary

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceConfigurationDeviceStateSummaryClient struct {
	Client *msgraph.Client
}

func NewDeviceConfigurationDeviceStateSummaryClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceConfigurationDeviceStateSummaryClient, error) {
	client, err := msgraph.NewClient(sdkApi, "deviceconfigurationdevicestatesummary", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceConfigurationDeviceStateSummaryClient: %+v", err)
	}

	return &DeviceConfigurationDeviceStateSummaryClient{
		Client: client,
	}, nil
}
