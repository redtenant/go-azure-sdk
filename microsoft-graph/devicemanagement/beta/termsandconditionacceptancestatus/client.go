package termsandconditionacceptancestatus

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TermsAndConditionAcceptanceStatusClient struct {
	Client *msgraph.Client
}

func NewTermsAndConditionAcceptanceStatusClientWithBaseURI(sdkApi sdkEnv.Api) (*TermsAndConditionAcceptanceStatusClient, error) {
	client, err := msgraph.NewClient(sdkApi, "termsandconditionacceptancestatus", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TermsAndConditionAcceptanceStatusClient: %+v", err)
	}

	return &TermsAndConditionAcceptanceStatusClient{
		Client: client,
	}, nil
}
