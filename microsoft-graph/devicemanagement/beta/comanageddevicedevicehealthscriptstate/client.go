package comanageddevicedevicehealthscriptstate

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComanagedDeviceDeviceHealthScriptStateClient struct {
	Client *msgraph.Client
}

func NewComanagedDeviceDeviceHealthScriptStateClientWithBaseURI(sdkApi sdkEnv.Api) (*ComanagedDeviceDeviceHealthScriptStateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "comanageddevicedevicehealthscriptstate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ComanagedDeviceDeviceHealthScriptStateClient: %+v", err)
	}

	return &ComanagedDeviceDeviceHealthScriptStateClient{
		Client: client,
	}, nil
}
