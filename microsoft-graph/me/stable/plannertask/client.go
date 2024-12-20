package plannertask

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerTaskClient struct {
	Client *msgraph.Client
}

func NewPlannerTaskClientWithBaseURI(sdkApi sdkEnv.Api) (*PlannerTaskClient, error) {
	client, err := msgraph.NewClient(sdkApi, "plannertask", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PlannerTaskClient: %+v", err)
	}

	return &PlannerTaskClient{
		Client: client,
	}, nil
}
