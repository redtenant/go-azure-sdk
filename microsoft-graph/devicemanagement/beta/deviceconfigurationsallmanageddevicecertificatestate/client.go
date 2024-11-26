package deviceconfigurationsallmanageddevicecertificatestate

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceConfigurationsAllManagedDeviceCertificateStateClient struct {
	Client *msgraph.Client
}

func NewDeviceConfigurationsAllManagedDeviceCertificateStateClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceConfigurationsAllManagedDeviceCertificateStateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "deviceconfigurationsallmanageddevicecertificatestate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceConfigurationsAllManagedDeviceCertificateStateClient: %+v", err)
	}

	return &DeviceConfigurationsAllManagedDeviceCertificateStateClient{
		Client: client,
	}, nil
}
