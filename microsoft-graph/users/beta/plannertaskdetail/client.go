package plannertaskdetail

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerTaskDetailClient struct {
	Client *msgraph.Client
}

func NewPlannerTaskDetailClientWithBaseURI(sdkApi sdkEnv.Api) (*PlannerTaskDetailClient, error) {
	client, err := msgraph.NewClient(sdkApi, "plannertaskdetail", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PlannerTaskDetailClient: %+v", err)
	}

	return &PlannerTaskDetailClient{
		Client: client,
	}, nil
}
