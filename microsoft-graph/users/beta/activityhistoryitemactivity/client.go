package activityhistoryitemactivity

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActivityHistoryItemActivityClient struct {
	Client *msgraph.Client
}

func NewActivityHistoryItemActivityClientWithBaseURI(sdkApi sdkEnv.Api) (*ActivityHistoryItemActivityClient, error) {
	client, err := msgraph.NewClient(sdkApi, "activityhistoryitemactivity", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ActivityHistoryItemActivityClient: %+v", err)
	}

	return &ActivityHistoryItemActivityClient{
		Client: client,
	}, nil
}
