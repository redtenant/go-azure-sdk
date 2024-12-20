package calendarcalendarviewexceptionoccurrencecalendar

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarCalendarViewExceptionOccurrenceCalendarClient struct {
	Client *msgraph.Client
}

func NewCalendarCalendarViewExceptionOccurrenceCalendarClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarCalendarViewExceptionOccurrenceCalendarClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendarcalendarviewexceptionoccurrencecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarCalendarViewExceptionOccurrenceCalendarClient: %+v", err)
	}

	return &CalendarCalendarViewExceptionOccurrenceCalendarClient{
		Client: client,
	}, nil
}
