package calendarviewexceptionoccurrenceinstance

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarViewExceptionOccurrenceInstanceClient struct {
	Client *msgraph.Client
}

func NewCalendarViewExceptionOccurrenceInstanceClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarViewExceptionOccurrenceInstanceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendarviewexceptionoccurrenceinstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarViewExceptionOccurrenceInstanceClient: %+v", err)
	}

	return &CalendarViewExceptionOccurrenceInstanceClient{
		Client: client,
	}, nil
}
