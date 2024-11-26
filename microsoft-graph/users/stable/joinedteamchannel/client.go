package joinedteamchannel

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinedTeamChannelClient struct {
	Client *msgraph.Client
}

func NewJoinedTeamChannelClientWithBaseURI(sdkApi sdkEnv.Api) (*JoinedTeamChannelClient, error) {
	client, err := msgraph.NewClient(sdkApi, "joinedteamchannel", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JoinedTeamChannelClient: %+v", err)
	}

	return &JoinedTeamChannelClient{
		Client: client,
	}, nil
}
