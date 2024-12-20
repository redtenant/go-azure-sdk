package joinedteamscheduleoffershiftrequest

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinedTeamScheduleOfferShiftRequestClient struct {
	Client *msgraph.Client
}

func NewJoinedTeamScheduleOfferShiftRequestClientWithBaseURI(sdkApi sdkEnv.Api) (*JoinedTeamScheduleOfferShiftRequestClient, error) {
	client, err := msgraph.NewClient(sdkApi, "joinedteamscheduleoffershiftrequest", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JoinedTeamScheduleOfferShiftRequestClient: %+v", err)
	}

	return &JoinedTeamScheduleOfferShiftRequestClient{
		Client: client,
	}, nil
}
