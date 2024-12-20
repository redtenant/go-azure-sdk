package teamworkinstalledappteamsappdefinition

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkInstalledAppTeamsAppDefinitionClient struct {
	Client *msgraph.Client
}

func NewTeamworkInstalledAppTeamsAppDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamworkInstalledAppTeamsAppDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamworkinstalledappteamsappdefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamworkInstalledAppTeamsAppDefinitionClient: %+v", err)
	}

	return &TeamworkInstalledAppTeamsAppDefinitionClient{
		Client: client,
	}, nil
}
