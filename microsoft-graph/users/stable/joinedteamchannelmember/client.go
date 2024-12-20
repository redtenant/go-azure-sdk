package joinedteamchannelmember

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinedTeamChannelMemberClient struct {
	Client *msgraph.Client
}

func NewJoinedTeamChannelMemberClientWithBaseURI(sdkApi sdkEnv.Api) (*JoinedTeamChannelMemberClient, error) {
	client, err := msgraph.NewClient(sdkApi, "joinedteamchannelmember", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JoinedTeamChannelMemberClient: %+v", err)
	}

	return &JoinedTeamChannelMemberClient{
		Client: client,
	}, nil
}
