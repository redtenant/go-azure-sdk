package deviceconfigurationuserstatesummary

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceConfigurationUserStateSummaryClient struct {
	Client *msgraph.Client
}

func NewDeviceConfigurationUserStateSummaryClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceConfigurationUserStateSummaryClient, error) {
	client, err := msgraph.NewClient(sdkApi, "deviceconfigurationuserstatesummary", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceConfigurationUserStateSummaryClient: %+v", err)
	}

	return &DeviceConfigurationUserStateSummaryClient{
		Client: client,
	}, nil
}
