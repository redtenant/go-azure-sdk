package driverootversioncontent

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveRootVersionContentClient struct {
	Client *msgraph.Client
}

func NewDriveRootVersionContentClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveRootVersionContentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driverootversioncontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveRootVersionContentClient: %+v", err)
	}

	return &DriveRootVersionContentClient{
		Client: client,
	}, nil
}
