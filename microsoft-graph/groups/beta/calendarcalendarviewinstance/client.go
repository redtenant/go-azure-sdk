package calendarcalendarviewinstance

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarCalendarViewInstanceClient struct {
	Client *msgraph.Client
}

func NewCalendarCalendarViewInstanceClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarCalendarViewInstanceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendarcalendarviewinstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarCalendarViewInstanceClient: %+v", err)
	}

	return &CalendarCalendarViewInstanceClient{
		Client: client,
	}, nil
}
