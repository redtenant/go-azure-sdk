package teamownermailboxsetting

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamOwnerMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewTeamOwnerMailboxSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamOwnerMailboxSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamownermailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamOwnerMailboxSettingClient: %+v", err)
	}

	return &TeamOwnerMailboxSettingClient{
		Client: client,
	}, nil
}
