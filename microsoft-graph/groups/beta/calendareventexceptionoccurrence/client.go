package calendareventexceptionoccurrence

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarEventExceptionOccurrenceClient struct {
	Client *msgraph.Client
}

func NewCalendarEventExceptionOccurrenceClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarEventExceptionOccurrenceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendareventexceptionoccurrence", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarEventExceptionOccurrenceClient: %+v", err)
	}

	return &CalendarEventExceptionOccurrenceClient{
		Client: client,
	}, nil
}
