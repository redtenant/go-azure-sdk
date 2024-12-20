package calendareventexceptionoccurrenceextension

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarEventExceptionOccurrenceExtensionClient struct {
	Client *msgraph.Client
}

func NewCalendarEventExceptionOccurrenceExtensionClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarEventExceptionOccurrenceExtensionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendareventexceptionoccurrenceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarEventExceptionOccurrenceExtensionClient: %+v", err)
	}

	return &CalendarEventExceptionOccurrenceExtensionClient{
		Client: client,
	}, nil
}
