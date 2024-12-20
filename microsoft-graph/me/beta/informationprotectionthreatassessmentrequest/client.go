package informationprotectionthreatassessmentrequest

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InformationProtectionThreatAssessmentRequestClient struct {
	Client *msgraph.Client
}

func NewInformationProtectionThreatAssessmentRequestClientWithBaseURI(sdkApi sdkEnv.Api) (*InformationProtectionThreatAssessmentRequestClient, error) {
	client, err := msgraph.NewClient(sdkApi, "informationprotectionthreatassessmentrequest", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating InformationProtectionThreatAssessmentRequestClient: %+v", err)
	}

	return &InformationProtectionThreatAssessmentRequestClient{
		Client: client,
	}, nil
}
