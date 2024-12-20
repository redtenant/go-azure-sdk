package manageddeviceencryptionstate

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceEncryptionStateClient struct {
	Client *msgraph.Client
}

func NewManagedDeviceEncryptionStateClientWithBaseURI(sdkApi sdkEnv.Api) (*ManagedDeviceEncryptionStateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "manageddeviceencryptionstate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedDeviceEncryptionStateClient: %+v", err)
	}

	return &ManagedDeviceEncryptionStateClient{
		Client: client,
	}, nil
}
