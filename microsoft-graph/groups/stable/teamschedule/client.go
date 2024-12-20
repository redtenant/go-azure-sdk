package teamschedule

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamScheduleClient struct {
	Client *msgraph.Client
}

func NewTeamScheduleClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamScheduleClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamschedule", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamScheduleClient: %+v", err)
	}

	return &TeamScheduleClient{
		Client: client,
	}, nil
}
