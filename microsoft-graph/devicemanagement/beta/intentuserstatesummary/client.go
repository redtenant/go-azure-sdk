package intentuserstatesummary

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntentUserStateSummaryClient struct {
	Client *msgraph.Client
}

func NewIntentUserStateSummaryClientWithBaseURI(sdkApi sdkEnv.Api) (*IntentUserStateSummaryClient, error) {
	client, err := msgraph.NewClient(sdkApi, "intentuserstatesummary", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IntentUserStateSummaryClient: %+v", err)
	}

	return &IntentUserStateSummaryClient{
		Client: client,
	}, nil
}
