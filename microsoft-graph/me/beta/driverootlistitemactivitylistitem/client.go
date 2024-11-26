package driverootlistitemactivitylistitem

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveRootListItemActivityListItemClient struct {
	Client *msgraph.Client
}

func NewDriveRootListItemActivityListItemClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveRootListItemActivityListItemClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driverootlistitemactivitylistitem", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveRootListItemActivityListItemClient: %+v", err)
	}

	return &DriveRootListItemActivityListItemClient{
		Client: client,
	}, nil
}
