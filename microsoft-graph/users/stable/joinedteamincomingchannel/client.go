package joinedteamincomingchannel

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinedTeamIncomingChannelClient struct {
	Client *msgraph.Client
}

func NewJoinedTeamIncomingChannelClientWithBaseURI(sdkApi sdkEnv.Api) (*JoinedTeamIncomingChannelClient, error) {
	client, err := msgraph.NewClient(sdkApi, "joinedteamincomingchannel", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JoinedTeamIncomingChannelClient: %+v", err)
	}

	return &JoinedTeamIncomingChannelClient{
		Client: client,
	}, nil
}
