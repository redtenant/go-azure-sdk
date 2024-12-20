package teamprimarychannelsharedwithteamallowedmember

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamPrimaryChannelSharedWithTeamAllowedMemberClient struct {
	Client *msgraph.Client
}

func NewTeamPrimaryChannelSharedWithTeamAllowedMemberClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamPrimaryChannelSharedWithTeamAllowedMemberClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamprimarychannelsharedwithteamallowedmember", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamPrimaryChannelSharedWithTeamAllowedMemberClient: %+v", err)
	}

	return &TeamPrimaryChannelSharedWithTeamAllowedMemberClient{
		Client: client,
	}, nil
}
