package eventexceptionoccurrencecalendar

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventExceptionOccurrenceCalendarClient struct {
	Client *msgraph.Client
}

func NewEventExceptionOccurrenceCalendarClientWithBaseURI(sdkApi sdkEnv.Api) (*EventExceptionOccurrenceCalendarClient, error) {
	client, err := msgraph.NewClient(sdkApi, "eventexceptionoccurrencecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EventExceptionOccurrenceCalendarClient: %+v", err)
	}

	return &EventExceptionOccurrenceCalendarClient{
		Client: client,
	}, nil
}
