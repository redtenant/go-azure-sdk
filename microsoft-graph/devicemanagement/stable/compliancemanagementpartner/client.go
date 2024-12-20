package compliancemanagementpartner

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComplianceManagementPartnerClient struct {
	Client *msgraph.Client
}

func NewComplianceManagementPartnerClientWithBaseURI(sdkApi sdkEnv.Api) (*ComplianceManagementPartnerClient, error) {
	client, err := msgraph.NewClient(sdkApi, "compliancemanagementpartner", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ComplianceManagementPartnerClient: %+v", err)
	}

	return &ComplianceManagementPartnerClient{
		Client: client,
	}, nil
}
