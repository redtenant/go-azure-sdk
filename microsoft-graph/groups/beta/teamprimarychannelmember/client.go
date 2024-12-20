package teamprimarychannelmember

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamPrimaryChannelMemberClient struct {
	Client *msgraph.Client
}

func NewTeamPrimaryChannelMemberClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamPrimaryChannelMemberClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamprimarychannelmember", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamPrimaryChannelMemberClient: %+v", err)
	}

	return &TeamPrimaryChannelMemberClient{
		Client: client,
	}, nil
}
