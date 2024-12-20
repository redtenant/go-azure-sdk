package conversationthreadpostinreplytoextension

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConversationThreadPostInReplyToExtensionClient struct {
	Client *msgraph.Client
}

func NewConversationThreadPostInReplyToExtensionClientWithBaseURI(sdkApi sdkEnv.Api) (*ConversationThreadPostInReplyToExtensionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "conversationthreadpostinreplytoextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ConversationThreadPostInReplyToExtensionClient: %+v", err)
	}

	return &ConversationThreadPostInReplyToExtensionClient{
		Client: client,
	}, nil
}
