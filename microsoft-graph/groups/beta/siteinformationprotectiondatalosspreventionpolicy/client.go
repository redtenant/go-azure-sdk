package siteinformationprotectiondatalosspreventionpolicy

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteInformationProtectionDataLossPreventionPolicyClient struct {
	Client *msgraph.Client
}

func NewSiteInformationProtectionDataLossPreventionPolicyClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteInformationProtectionDataLossPreventionPolicyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteinformationprotectiondatalosspreventionpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteInformationProtectionDataLossPreventionPolicyClient: %+v", err)
	}

	return &SiteInformationProtectionDataLossPreventionPolicyClient{
		Client: client,
	}, nil
}
