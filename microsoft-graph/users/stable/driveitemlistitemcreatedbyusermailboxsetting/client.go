package driveitemlistitemcreatedbyusermailboxsetting

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItemListItemCreatedByUserMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewDriveItemListItemCreatedByUserMailboxSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveItemListItemCreatedByUserMailboxSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driveitemlistitemcreatedbyusermailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveItemListItemCreatedByUserMailboxSettingClient: %+v", err)
	}

	return &DriveItemListItemCreatedByUserMailboxSettingClient{
		Client: client,
	}, nil
}
