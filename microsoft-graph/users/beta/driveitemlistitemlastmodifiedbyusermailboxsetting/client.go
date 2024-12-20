package driveitemlistitemlastmodifiedbyusermailboxsetting

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItemListItemLastModifiedByUserMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewDriveItemListItemLastModifiedByUserMailboxSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveItemListItemLastModifiedByUserMailboxSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driveitemlistitemlastmodifiedbyusermailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveItemListItemLastModifiedByUserMailboxSettingClient: %+v", err)
	}

	return &DriveItemListItemLastModifiedByUserMailboxSettingClient{
		Client: client,
	}, nil
}
