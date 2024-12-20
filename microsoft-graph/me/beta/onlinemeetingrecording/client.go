package onlinemeetingrecording

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnlineMeetingRecordingClient struct {
	Client *msgraph.Client
}

func NewOnlineMeetingRecordingClientWithBaseURI(sdkApi sdkEnv.Api) (*OnlineMeetingRecordingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onlinemeetingrecording", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnlineMeetingRecordingClient: %+v", err)
	}

	return &OnlineMeetingRecordingClient{
		Client: client,
	}, nil
}
