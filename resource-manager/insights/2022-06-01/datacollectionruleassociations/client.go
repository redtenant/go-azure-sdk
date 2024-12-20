package datacollectionruleassociations

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataCollectionRuleAssociationsClient struct {
	Client *resourcemanager.Client
}

func NewDataCollectionRuleAssociationsClientWithBaseURI(sdkApi sdkEnv.Api) (*DataCollectionRuleAssociationsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "datacollectionruleassociations", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DataCollectionRuleAssociationsClient: %+v", err)
	}

	return &DataCollectionRuleAssociationsClient{
		Client: client,
	}, nil
}
