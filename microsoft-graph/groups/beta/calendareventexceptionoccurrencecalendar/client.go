package calendareventexceptionoccurrencecalendar

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarEventExceptionOccurrenceCalendarClient struct {
	Client *msgraph.Client
}

func NewCalendarEventExceptionOccurrenceCalendarClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarEventExceptionOccurrenceCalendarClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendareventexceptionoccurrencecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarEventExceptionOccurrenceCalendarClient: %+v", err)
	}

	return &CalendarEventExceptionOccurrenceCalendarClient{
		Client: client,
	}, nil
}
