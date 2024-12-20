package drivelistpermission

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveListPermissionClient struct {
	Client *msgraph.Client
}

func NewDriveListPermissionClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveListPermissionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "drivelistpermission", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveListPermissionClient: %+v", err)
	}

	return &DriveListPermissionClient{
		Client: client,
	}, nil
}
