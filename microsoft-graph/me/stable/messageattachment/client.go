package messageattachment

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MessageAttachmentClient struct {
	Client *msgraph.Client
}

func NewMessageAttachmentClientWithBaseURI(sdkApi sdkEnv.Api) (*MessageAttachmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "messageattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MessageAttachmentClient: %+v", err)
	}

	return &MessageAttachmentClient{
		Client: client,
	}, nil
}
