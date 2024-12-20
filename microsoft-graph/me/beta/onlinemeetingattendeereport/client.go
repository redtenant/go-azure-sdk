package onlinemeetingattendeereport

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnlineMeetingAttendeeReportClient struct {
	Client *msgraph.Client
}

func NewOnlineMeetingAttendeeReportClientWithBaseURI(sdkApi sdkEnv.Api) (*OnlineMeetingAttendeeReportClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onlinemeetingattendeereport", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnlineMeetingAttendeeReportClient: %+v", err)
	}

	return &OnlineMeetingAttendeeReportClient{
		Client: client,
	}, nil
}
