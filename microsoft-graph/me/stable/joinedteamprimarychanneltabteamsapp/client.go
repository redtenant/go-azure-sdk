package joinedteamprimarychanneltabteamsapp

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinedTeamPrimaryChannelTabTeamsAppClient struct {
	Client *msgraph.Client
}

func NewJoinedTeamPrimaryChannelTabTeamsAppClientWithBaseURI(sdkApi sdkEnv.Api) (*JoinedTeamPrimaryChannelTabTeamsAppClient, error) {
	client, err := msgraph.NewClient(sdkApi, "joinedteamprimarychanneltabteamsapp", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JoinedTeamPrimaryChannelTabTeamsAppClient: %+v", err)
	}

	return &JoinedTeamPrimaryChannelTabTeamsAppClient{
		Client: client,
	}, nil
}
