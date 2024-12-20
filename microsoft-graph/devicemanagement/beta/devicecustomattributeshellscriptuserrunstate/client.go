package devicecustomattributeshellscriptuserrunstate

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCustomAttributeShellScriptUserRunStateClient struct {
	Client *msgraph.Client
}

func NewDeviceCustomAttributeShellScriptUserRunStateClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceCustomAttributeShellScriptUserRunStateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicecustomattributeshellscriptuserrunstate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceCustomAttributeShellScriptUserRunStateClient: %+v", err)
	}

	return &DeviceCustomAttributeShellScriptUserRunStateClient{
		Client: client,
	}, nil
}
