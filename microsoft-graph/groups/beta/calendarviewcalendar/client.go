package calendarviewcalendar

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarViewCalendarClient struct {
	Client *msgraph.Client
}

func NewCalendarViewCalendarClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarViewCalendarClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendarviewcalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarViewCalendarClient: %+v", err)
	}

	return &CalendarViewCalendarClient{
		Client: client,
	}, nil
}
