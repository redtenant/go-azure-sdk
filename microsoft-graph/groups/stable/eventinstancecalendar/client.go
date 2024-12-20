package eventinstancecalendar

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventInstanceCalendarClient struct {
	Client *msgraph.Client
}

func NewEventInstanceCalendarClientWithBaseURI(sdkApi sdkEnv.Api) (*EventInstanceCalendarClient, error) {
	client, err := msgraph.NewClient(sdkApi, "eventinstancecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EventInstanceCalendarClient: %+v", err)
	}

	return &EventInstanceCalendarClient{
		Client: client,
	}, nil
}
