package drivelistcolumn

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveListColumnClient struct {
	Client *msgraph.Client
}

func NewDriveListColumnClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveListColumnClient, error) {
	client, err := msgraph.NewClient(sdkApi, "drivelistcolumn", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveListColumnClient: %+v", err)
	}

	return &DriveListColumnClient{
		Client: client,
	}, nil
}
