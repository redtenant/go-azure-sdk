package plannerrosterplan

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerRosterPlanClient struct {
	Client *msgraph.Client
}

func NewPlannerRosterPlanClientWithBaseURI(sdkApi sdkEnv.Api) (*PlannerRosterPlanClient, error) {
	client, err := msgraph.NewClient(sdkApi, "plannerrosterplan", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PlannerRosterPlanClient: %+v", err)
	}

	return &PlannerRosterPlanClient{
		Client: client,
	}, nil
}
