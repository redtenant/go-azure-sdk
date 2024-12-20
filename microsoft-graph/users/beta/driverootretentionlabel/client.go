package driverootretentionlabel

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveRootRetentionLabelClient struct {
	Client *msgraph.Client
}

func NewDriveRootRetentionLabelClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveRootRetentionLabelClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driverootretentionlabel", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveRootRetentionLabelClient: %+v", err)
	}

	return &DriveRootRetentionLabelClient{
		Client: client,
	}, nil
}
