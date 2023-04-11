package budgets

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BudgetsClient struct {
	Client *resourcemanager.Client
}

func NewBudgetsClientWithBaseURI(api environments.Api) (*BudgetsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "budgets", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating BudgetsClient: %+v", err)
	}

	return &BudgetsClient{
		Client: client,
	}, nil
}
