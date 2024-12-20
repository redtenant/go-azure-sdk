package calendarviewattachment

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarViewAttachmentClient struct {
	Client *msgraph.Client
}

func NewCalendarViewAttachmentClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarViewAttachmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendarviewattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarViewAttachmentClient: %+v", err)
	}

	return &CalendarViewAttachmentClient{
		Client: client,
	}, nil
}
