package devicecompliancepolicysettingstatesummary

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCompliancePolicySettingStateSummaryClient struct {
	Client *msgraph.Client
}

func NewDeviceCompliancePolicySettingStateSummaryClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceCompliancePolicySettingStateSummaryClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicecompliancepolicysettingstatesummary", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceCompliancePolicySettingStateSummaryClient: %+v", err)
	}

	return &DeviceCompliancePolicySettingStateSummaryClient{
		Client: client,
	}, nil
}
