package teaminstalledappteamsapp

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamInstalledAppTeamsAppClient struct {
	Client *msgraph.Client
}

func NewTeamInstalledAppTeamsAppClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamInstalledAppTeamsAppClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teaminstalledappteamsapp", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamInstalledAppTeamsAppClient: %+v", err)
	}

	return &TeamInstalledAppTeamsAppClient{
		Client: client,
	}, nil
}
