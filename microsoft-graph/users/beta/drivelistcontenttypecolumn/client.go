package drivelistcontenttypecolumn

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveListContentTypeColumnClient struct {
	Client *msgraph.Client
}

func NewDriveListContentTypeColumnClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveListContentTypeColumnClient, error) {
	client, err := msgraph.NewClient(sdkApi, "drivelistcontenttypecolumn", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveListContentTypeColumnClient: %+v", err)
	}

	return &DriveListContentTypeColumnClient{
		Client: client,
	}, nil
}
