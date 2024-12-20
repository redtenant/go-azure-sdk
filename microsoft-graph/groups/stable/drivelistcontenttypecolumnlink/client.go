package drivelistcontenttypecolumnlink

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveListContentTypeColumnLinkClient struct {
	Client *msgraph.Client
}

func NewDriveListContentTypeColumnLinkClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveListContentTypeColumnLinkClient, error) {
	client, err := msgraph.NewClient(sdkApi, "drivelistcontenttypecolumnlink", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveListContentTypeColumnLinkClient: %+v", err)
	}

	return &DriveListContentTypeColumnLinkClient{
		Client: client,
	}, nil
}
