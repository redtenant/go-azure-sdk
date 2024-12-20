package driverootchild

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveRootChildClient struct {
	Client *msgraph.Client
}

func NewDriveRootChildClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveRootChildClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driverootchild", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveRootChildClient: %+v", err)
	}

	return &DriveRootChildClient{
		Client: client,
	}, nil
}
