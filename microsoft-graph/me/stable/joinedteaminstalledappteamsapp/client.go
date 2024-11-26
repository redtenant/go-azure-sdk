package joinedteaminstalledappteamsapp

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinedTeamInstalledAppTeamsAppClient struct {
	Client *msgraph.Client
}

func NewJoinedTeamInstalledAppTeamsAppClientWithBaseURI(sdkApi sdkEnv.Api) (*JoinedTeamInstalledAppTeamsAppClient, error) {
	client, err := msgraph.NewClient(sdkApi, "joinedteaminstalledappteamsapp", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JoinedTeamInstalledAppTeamsAppClient: %+v", err)
	}

	return &JoinedTeamInstalledAppTeamsAppClient{
		Client: client,
	}, nil
}
