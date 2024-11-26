package importeddeviceidentity

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImportedDeviceIdentityClient struct {
	Client *msgraph.Client
}

func NewImportedDeviceIdentityClientWithBaseURI(sdkApi sdkEnv.Api) (*ImportedDeviceIdentityClient, error) {
	client, err := msgraph.NewClient(sdkApi, "importeddeviceidentity", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ImportedDeviceIdentityClient: %+v", err)
	}

	return &ImportedDeviceIdentityClient{
		Client: client,
	}, nil
}
