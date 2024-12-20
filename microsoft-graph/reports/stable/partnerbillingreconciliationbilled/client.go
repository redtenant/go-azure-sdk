package partnerbillingreconciliationbilled

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PartnerBillingReconciliationBilledClient struct {
	Client *msgraph.Client
}

func NewPartnerBillingReconciliationBilledClientWithBaseURI(sdkApi sdkEnv.Api) (*PartnerBillingReconciliationBilledClient, error) {
	client, err := msgraph.NewClient(sdkApi, "partnerbillingreconciliationbilled", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PartnerBillingReconciliationBilledClient: %+v", err)
	}

	return &PartnerBillingReconciliationBilledClient{
		Client: client,
	}, nil
}
