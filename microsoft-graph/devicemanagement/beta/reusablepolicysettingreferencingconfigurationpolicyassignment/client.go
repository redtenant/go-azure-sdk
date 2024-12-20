package reusablepolicysettingreferencingconfigurationpolicyassignment

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReusablePolicySettingReferencingConfigurationPolicyAssignmentClient struct {
	Client *msgraph.Client
}

func NewReusablePolicySettingReferencingConfigurationPolicyAssignmentClientWithBaseURI(sdkApi sdkEnv.Api) (*ReusablePolicySettingReferencingConfigurationPolicyAssignmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "reusablepolicysettingreferencingconfigurationpolicyassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ReusablePolicySettingReferencingConfigurationPolicyAssignmentClient: %+v", err)
	}

	return &ReusablePolicySettingReferencingConfigurationPolicyAssignmentClient{
		Client: client,
	}, nil
}
