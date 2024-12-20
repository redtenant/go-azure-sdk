package calendarviewexceptionoccurrencecalendar

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarViewExceptionOccurrenceCalendarClient struct {
	Client *msgraph.Client
}

func NewCalendarViewExceptionOccurrenceCalendarClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarViewExceptionOccurrenceCalendarClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendarviewexceptionoccurrencecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarViewExceptionOccurrenceCalendarClient: %+v", err)
	}

	return &CalendarViewExceptionOccurrenceCalendarClient{
		Client: client,
	}, nil
}
