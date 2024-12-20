package joinedteam

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinedTeamClient struct {
	Client *msgraph.Client
}

func NewJoinedTeamClientWithBaseURI(sdkApi sdkEnv.Api) (*JoinedTeamClient, error) {
	client, err := msgraph.NewClient(sdkApi, "joinedteam", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JoinedTeamClient: %+v", err)
	}

	return &JoinedTeamClient{
		Client: client,
	}, nil
}
