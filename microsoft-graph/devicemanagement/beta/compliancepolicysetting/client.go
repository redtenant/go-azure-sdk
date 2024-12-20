package compliancepolicysetting

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CompliancePolicySettingClient struct {
	Client *msgraph.Client
}

func NewCompliancePolicySettingClientWithBaseURI(sdkApi sdkEnv.Api) (*CompliancePolicySettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "compliancepolicysetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CompliancePolicySettingClient: %+v", err)
	}

	return &CompliancePolicySettingClient{
		Client: client,
	}, nil
}
