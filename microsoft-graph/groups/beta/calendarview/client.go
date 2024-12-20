package calendarview

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarViewClient struct {
	Client *msgraph.Client
}

func NewCalendarViewClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarViewClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendarview", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarViewClient: %+v", err)
	}

	return &CalendarViewClient{
		Client: client,
	}, nil
}
