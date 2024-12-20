package eventinstanceexceptionoccurrenceattachment

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventInstanceExceptionOccurrenceAttachmentClient struct {
	Client *msgraph.Client
}

func NewEventInstanceExceptionOccurrenceAttachmentClientWithBaseURI(sdkApi sdkEnv.Api) (*EventInstanceExceptionOccurrenceAttachmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "eventinstanceexceptionoccurrenceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EventInstanceExceptionOccurrenceAttachmentClient: %+v", err)
	}

	return &EventInstanceExceptionOccurrenceAttachmentClient{
		Client: client,
	}, nil
}
