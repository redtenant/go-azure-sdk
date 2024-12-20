package comanageddevicelogcollectionrequest

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComanagedDeviceLogCollectionRequestClient struct {
	Client *msgraph.Client
}

func NewComanagedDeviceLogCollectionRequestClientWithBaseURI(sdkApi sdkEnv.Api) (*ComanagedDeviceLogCollectionRequestClient, error) {
	client, err := msgraph.NewClient(sdkApi, "comanageddevicelogcollectionrequest", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ComanagedDeviceLogCollectionRequestClient: %+v", err)
	}

	return &ComanagedDeviceLogCollectionRequestClient{
		Client: client,
	}, nil
}
