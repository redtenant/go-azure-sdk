package deviceenrollmentconfigurationassignment

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceEnrollmentConfigurationAssignmentClient struct {
	Client *msgraph.Client
}

func NewDeviceEnrollmentConfigurationAssignmentClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceEnrollmentConfigurationAssignmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "deviceenrollmentconfigurationassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceEnrollmentConfigurationAssignmentClient: %+v", err)
	}

	return &DeviceEnrollmentConfigurationAssignmentClient{
		Client: client,
	}, nil
}
