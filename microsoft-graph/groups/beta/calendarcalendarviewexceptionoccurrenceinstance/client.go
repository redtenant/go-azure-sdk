package calendarcalendarviewexceptionoccurrenceinstance

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarCalendarViewExceptionOccurrenceInstanceClient struct {
	Client *msgraph.Client
}

func NewCalendarCalendarViewExceptionOccurrenceInstanceClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarCalendarViewExceptionOccurrenceInstanceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendarcalendarviewexceptionoccurrenceinstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarCalendarViewExceptionOccurrenceInstanceClient: %+v", err)
	}

	return &CalendarCalendarViewExceptionOccurrenceInstanceClient{
		Client: client,
	}, nil
}
